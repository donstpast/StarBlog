package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starblog/utils/errmsg"
)

func FileUploadLocal(c *gin.Context) {
	// 从请求中获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 修改保存文件的路径
	filePath := "./upload/" + file.Filename
	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    "/upload/" + file.Filename,
		"message": "File uploaded successfully",
	})
}
