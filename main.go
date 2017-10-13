package main

import (
	"sun/controller"
	"sun/core"
	"sun/mds"
	"sun/model"
)

func main() {
	// 数据库初始化
	model.Init()

	f := core.Default()
	// 处理静态资源路径
	f.Use(mds.Login())
	f.Use(mds.Static("/static"))
	f.Use(mds.Favicon("/static/favicon.ico"))

	// 添加路由
	routerMd, routers := mds.SRouter()

	routers.Add("/", &controller.Index{})
	routers.Add("/list", &controller.List{})
	routers.Add("/publish", &controller.Publish{})
	routers.Add("/detail", &controller.Detail{})
	routers.Add("/404", &controller.NotFound{})

	routers.Add("/api/post", &controller.SJSON{})
	routers.Add("/api/comments", &controller.Comments{})
	routers.Add("/api/comments/add", &controller.Comments{})
	routers.Add("/api/delete/redis", &controller.DeleteRedis{})
	routers.NotMatch("/404") // 404页面处理
	f.Use(routerMd)

	f.Listen("8965")
}
