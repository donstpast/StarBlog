package errmsg

//错误信息处理

// 定义一些错误常量（最常见的时候状态码）
const (
	SUCCESS = 200
	ERROR   = 500

	// 之后定义一些错误代码，我们规定如下:

	//code在(1000,2000)为用户模块错误:
	ERROR_USER_NOT_EXIST     = 1001 //用户不存在
	ERROR_USERNAME_USED      = 1002 //用户名已被使用
	ERROR_PASSWORD_WRONG     = 1003 //密码错误
	ERROR_TOKEN_NOT_EXIST    = 1004 //token不存在
	ERROT_TOKEN_OVERTIME     = 1005 //token过期
	ERROT_TOKEN_WRONG        = 1006 //token错误
	ERROR_TOKEN_FORMAT_WRONG = 1007 //token格式错误
	//code在(2000,3000)为文章模块错误
	//code在(3000,4000)为分类模块错误
)

// 定义一个map
// int处为上方状态码常量，string处为需要抛出的错误信息
var codeMsg = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "FAIL",
	ERROR_USER_NOT_EXIST:     "用户不存在",
	ERROR_USERNAME_USED:      "用户名已被占用",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_TOKEN_NOT_EXIST:    "TOKEN不存在",
	ERROT_TOKEN_OVERTIME:     "TOKEN已过期",
	ERROT_TOKEN_WRONG:        "TOKEN错误",
	ERROR_TOKEN_FORMAT_WRONG: "TOKEN格式错误",
}

// GetErrMsg 用来输出错误信息的函数,传入一个int型，返回一个string类型
func GetErrMsg(code int) string {
	return codeMsg[code]
}
