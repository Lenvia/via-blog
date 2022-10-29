package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"via-blog/middleware"
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
		setToken(c, formData)  // 生成并设置token
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
	j:= middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer: "viaBlog",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token": token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": user.Username,
		"id": user.ID,
		"message": errmsg.GetErrMsg(200),
		"token": token,
	})
	return

}