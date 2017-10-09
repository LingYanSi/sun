package controller

import (
	"fmt"
	"sun/core"
)

// Detail 处理主页
type Detail struct {
	core.Base
}

// GET 处理get请求
func (c *Detail) GET() {
	// 允许跨域请求
	c.Allow()
	// c.SetHeader("Content-Type", "text/html; charset=utf8")
	// c.Res.WriteHeader(200)

	name, err := c.Redis.Get("name").Result()
	if err != nil {
		name = "出错了"
		fmt.Println("错误信息: ", err)
	}
	fmt.Println("redis name:", name)
	c.HTML("Detail", core.J{
		"title": "详情页",
		"body":  "人才啊" + name,
	})
}
