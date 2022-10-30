package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"via-blog/model"
	"via-blog/utils/errmsg"
)

// Upload 上传图片接口
func Upload(c *gin.Context)  {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size

	url, code := model.UploadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
		"url": url,
	})
}