package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// Logger 自定义log中间件
func Logger() gin.HandlerFunc {
	logPath := "log/log"                                          // 定义日志文件路径
	linkPath := "latest_log.log"                                  // 定义日志软链接文件路径
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755) // 打开日志文件，如果不存在则创建，使用os.OpenFile方法，文件权限为0755
	if err != nil {                                               // 如果打开或创建文件出错，输出错误信息
		fmt.Println(err)
	}
	// 实例化一个log对象
	logger := log.New()

	// 设置日志输出格式为JSON格式
	log.SetFormatter(&log.JSONFormatter{})
	logger.Out = src // 设置日志的输出文件为src

	// 设置log的级别为DebugLevel
	logger.SetLevel(log.DebugLevel)

	// 设置log的写入器，使用rotalog来进行日志的分割和清理
	logWriter, _ := rotalog.New(
		logPath+"%Y%m%d.log",
		rotalog.WithMaxAge(7*24*time.Hour),     //日志最多保存7天，之后删除
		rotalog.WithRotationTime(24*time.Hour), //每24小时进行一次分割
		rotalog.WithLinkName(linkPath),         //软链接文件
	)
	writeMap := lfshook.WriterMap{ // 构建log的写入器映射表
		log.InfoLevel:  logWriter, // Info级别的log使用logWriter写入
		log.FatalLevel: logWriter, // Fatal级别的log使用logWriter写入
		log.TraceLevel: logWriter, // Trace级别的log使用logWriter写入
		log.DebugLevel: logWriter, // Debug级别的log使用logWriter写入
		log.WarnLevel:  logWriter, // Warn级别的log使用logWriter写入
		log.ErrorLevel: logWriter, // Error级别的log使用logWriter写入
		log.PanicLevel: logWriter, // Panic级别的log使用logWriter写入
	}
	Hook := lfshook.NewHook(writeMap, &log.TextFormatter{ // 构建log的钩子对象，将日志写入logWriter中
		TimestampFormat: "2006-01-02 15:04:05", // 时间戳格式
	})
	logger.AddHook(Hook) // 添加钩子对象到logger中

	return func(c *gin.Context) { // 返回一个gin的处理函数
		startTime := time.Now() // 记录处理开始的时间

		c.Next() // 调用gin.Context的Next方法，处理后续的中间件和请求处理函数

		stopTime := time.Since(startTime)                          // 计算处理耗时
		spendTime := fmt.Sprintf("%d ms", stopTime.Milliseconds()) // 耗时转换成字符串格式，单位为毫秒

		hostName, err := os.Hostname() // 获取本机hostname
		if err != nil {                // 如果获取hostname出错，将其设置为unknown
			hostName = "unknown"
		}

		statusCode := c.Writer.Status()    // 获取HTTP响应状态码
		clientIp := c.ClientIP()           // 获取客户端IP地址
		userAgent := c.Request.UserAgent() // 获取客户端User-Agent头信息
		dataSize := c.Writer.Size()        // 获取HTTP响应正文长度
		if dataSize < 0 {                  // 如果获取dataSize小于0，则代表有错，将其设置为0
			dataSize = 0
		}

		reqMethod := c.Request.Method   // 获取HTTP请求方法
		reqPath := c.Request.RequestURI // 获取HTTP请求URI
		// 构建log.Entry对象，包含HTTP请求的各种信息
		entry := logger.WithFields(log.Fields{
			"RequestPath":   reqPath,
			"RequestMethod": reqMethod,
			"ClientIP":      clientIp,
			"HostName":      hostName,
			"Status":        statusCode,
			"SpendTime":     spendTime,
			"DataSize":      dataSize,
			"Agent":         userAgent,
		})
		// 如果处理过程中出现错误，将错误信息记录到log中
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}

		if statusCode >= 500 { // 如果HTTP响应状态码大于等于500，记录Error级别的log
			entry.Error()
		} else if statusCode >= 400 { // 如果HTTP响应状态码大于等于400，记录Warn级别的log
			entry.Warn()
		} else { // 如果HTTP响应状态码小于400，记录Info级别的log
			entry.Info()
		}
	}
}
