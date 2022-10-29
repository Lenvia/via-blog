package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"via-blog/model"
	"via-blog/utils/errmsg"
)

// Login 后台登录
func Login(c *gin.Context)  {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)

	var token string
	var code int
	// 返回的formData包含加密后的真实密码，下面c.JSON不传
	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == errmsg.SUCCESS{
		// todo 生成 token
	} else{
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data": formData.Username,
			"id": formData.ID,
			"message": errmsg.GetErrMsg(code),
			"token": token,
		})
	}
}

// LoginFront 前台登录（不用token）
func LoginFront(c *gin.Context)  {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)

	var code int
	formData, code = model.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": formData.Username,
		"id": formData.ID,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo token生成函数
func setToken(c *gin.Context, user model.User)  {

}