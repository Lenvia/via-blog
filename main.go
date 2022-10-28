package main

import (
	"viaBlog/model"
	"viaBlog/routers"
)

func main()  {
	// 引入数据库
	model.InitDb()
	// 引入路由组件
	routers.InitRouter()
}