package routers

import (
	"github.com/gin-gonic/gin"
	"via-blog/api/v1"
	"via-blog/middleware"
	"via-blog/utils"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)

	r:= gin.Default()
	//r := gin.New()

	//r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	/*
		后台管理路由接口
	 */
	// 注意：authR 不是区分身份，而是说访问 authR路由下的方法，需要通过额外的中间件验证
	authR := r.Group("api/v1")
	// 每个用户登录 都会走另一个分组 router 的登录控制，然后获得token，只有带着token才能访问 authR 分组下的方法
	authR.Use(middleware.JwtToken())
	{
		// 用户模块的接口
		authR.GET("admin/user", v1.GetUsers)
		authR.PUT("user/:id", v1.UpdateUser)
		authR.DELETE("user/:id", v1.DeleteUser)
		authR.PUT("admin/changepwd/:id", v1.ChangeUserPassWord) // 管理员才能修改密码

		// 分类模块的接口
		authR.GET("admin/category", v1.GetCategories)
		authR.POST("category/add", v1.AddCategory)
		authR.PUT("category/:id", v1.UpdateCategory)
		authR.DELETE("category/:id", v1.DeleteCategory)


		// 文章模块的接口
		authR.GET("admin/article/:id", v1.GetArticle)
		authR.GET("admin/article", v1.GetArticles)
		authR.POST("article/add", v1.AddArticle)
		authR.PUT("article/:id", v1.UpdateArticle)
		authR.DELETE("article/id", v1.DeleteArticle)

		authR.POST("upload", v1.Upload)// 上传文件

		// 更新个人设置
		authR.GET("admin/profile/:id", v1.GetProfile)
		authR.PUT("profile/:id", v1.UpdateProfile)

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
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetProfile)

	}

	_ = r.Run(utils.HttpPort)
}
