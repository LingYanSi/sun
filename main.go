package main

import (
	"sun/controller"
	"sun/core"
)

func main() {
	f := core.Default()
	f.Add("/", &controller.Index{})
	f.Add("/404", &controller.NotFound{})
	f.Add("/my", &controller.Index{})
	f.Add("/api/post", &controller.SJSON{})
	f.Add("/api/comments", &controller.Comments{})
	f.Add("/api/comments/add", &controller.Comments{})
	f.Add("/api/delete/redis", &controller.DeleteRedis{})
	f.NotMatch("/404")
	f.Listen("8965")
}

// package main // 包名

// import (
// 	// 引用controller
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	"github.com/thinkerou/favicon"
// )

// func main() {
// 	// gin.SetMode(gin.ReleaseMode)
// 	r := gin.Default()

// 	// 静态文件
// 	r.Static("/assets", "./static")

// 	// icon
// 	// r.StaticFile("/favicon.ico", "./static/favicon.ico")

// 	r.Use(favicon.New("./favicon.ico"))

// 	// string
// 	// r.GET("/my", func(c *gin.Context) {
// 	// 	h := controller.Home{controller.Base{Req: c.Request, Res: c.Writer}}
// 	// 	h.Get()
// 	// })

// 	r.GET("/str", func(c *gin.Context) {
// 		c.String(http.StatusOK, "hello go")
// 	})

// 	// json
// 	r.GET("/json", func(c *gin.Context) {
// 		names := []string{"lena", "austin", "foo"}
// 		// Will output  :   while(1);["lena","austin","foo"]
// 		c.JSON(http.StatusOK, names)
// 	})

// 	// html render
// 	r.LoadHTMLGlob("views/*")
// 	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
// 	r.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.html", gin.H{
// 			"title": "Main website",
// 			"body":  "hello go",
// 		})
// 	})

// 	// 重定向
// 	r.GET("/redirect", func(c *gin.Context) {
// 		c.Redirect(http.StatusMovedPermanently, "/")
// 	})

// 	// 聊天系统
// 	r.GET("/room/:roomid", roomGET)
// 	r.POST("/room/:roomid", roomPOST)
// 	r.DELETE("/room/:roomid", roomDELETE)
// 	r.GET("/stream/:roomid", stream)

// 	// 路由不匹配
// 	r.NoRoute(func(c *gin.Context) {
// 		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
// 	})

// 	r.Run(":8965") // listen and serve on 0.0.0.0:8080
// }
