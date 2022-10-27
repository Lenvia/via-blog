package routers

import (
	"github.com/gin-gonic/gin"
	"viaBlog/utils"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	//r := gin.New()
	r:= gin.Default()


	/*
		后台管理路由接口
	 */
	// 管理员router，需要加中间件检查权限
	auth_r := r.Group("api/v1")
	{
		// 用户模块的接口
		// 分类模块的接口
		// 文章模块的接口

	}

	router := r.Group("api/v1")
	{
		// 用户信息模块
		// 分类模块
		// 文章模块
		// 登录控制
		
	}
}
