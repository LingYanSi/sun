package main // 包名

import (
	"fmt"
	"net/http"
	"sun/controller" // 引用controller

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 静态文件
	r.Static("/assets", "./static")

	// icon
	// r.StaticFile("/favicon.ico", "./static/favicon.ico")

	// r.Use(favicon.New("./static/favicon.ico"))

	// string
	// r.GET("/my", func(c *gin.Context) {
	// 	h := controller.Home{controller.Base{Req: c.Request, Res: c.Writer}}
	// 	h.Get()
	// })

	r.GET("/my1", func(c *gin.Context) {
		// h := controller.Home{controller.Base{Req: c.Request, Res: c.Writer}}
		f := controller.Sun{}
		f.Init()
		f.AddRouter("/my1", &controller.Home{})
		f.HandleReq(c.Request, c.Writer)
	})

	r.GET("/str", func(c *gin.Context) {
		c.String(http.StatusOK, "hello go")
	})

	// json
	r.GET("/json", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// Will output  :   while(1);["lena","austin","foo"]
		c.JSON(http.StatusOK, names)
	})

	// html render
	r.LoadHTMLGlob("views/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"body":  "hello go",
		})
	})

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.GET("/room/:roomid", roomGET)
	r.POST("/room/:roomid", roomPOST)
	r.DELETE("/room/:roomid", roomDELETE)
	r.GET("/stream/:roomid", stream)

	// 路由不匹配
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run(":8965") // listen and serve on 0.0.0.0:8080
	fmt.Println("http://localhost:8965")
}
