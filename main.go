package main

import (
	"via-blog/dao"
	"via-blog/routers"
)

func main()  {
	// 引入数据库
	dao.InitDb()
	// 引入路由组件
	routers.InitRouter()
}