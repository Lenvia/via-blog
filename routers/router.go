package routers

import (
	"github.com/gin-gonic/gin"
	v1 "viaBlog/api/v1"
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
	authR := r.Group("api/v1")
	{
		// 用户模块的接口
		authR.GET("admin/user", v1.GetUsers)
		authR.PUT("user/:id", v1.UpdateUser)
		authR.DELETE("user/:id", v1.DeleteUser)
		authR.PUT("admin/changepwd/:id", v1.ChangeUserPassWord) // 管理员才能修改密码

		// 分类模块的接口
		authR.GET("admin/category", v1.GetCategory)
		authR.POST("add", v1.AddCategory)
		authR.PUT("category/:id", v1.UpdateCategory)
		authR.DELETE("category/:id", v1.DeleteCategory)


		// 文章模块的接口
		authR.GET("admin/article/:id", v1.GetArticle)
		authR.GET("admin/article", v1.GetArticles)
		authR.POST("article/add", v1.AddArticle)
		authR.PUT("article/:id", v1.UpdateArticle)
		authR.DELETE("article/id", v1.DeleteArticle)

	}

	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)  // 任何人都能注册，角色已被限制非管理员
		router.GET("user/:id", v1.GetUser)
		router.GET("user", v1.GetUsers)

		// 分类模块
		router.GET("category", v1.GetCategories)  // 查看所有分类
		router.GET("category/:id", v1.GetCategory)

		// 文章模块
		router.GET("article", v1.GetArticles)
		router.GET("article/list/:id", v1.GetCateArticles)  // 感觉这里改成 :cid 更好
		router.GET("article/:id", v1.GetArticle)
		// 登录控制
		
	}

	_ = r.Run(utils.HttpPort)
}
