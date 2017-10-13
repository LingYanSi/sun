package controller

import (
	"fmt"
	"sun/core"
	"sun/model"
)

// Detail 处理主页
type Detail struct {
	core.Base
}

// GET 处理get请求
func (c *Detail) GET() {
	// 允许跨域请求
	// c.Allow()
	// c.SetHeader("Content-Type", "text/html; charset=utf8")
	// c.Res.WriteHeader(200)
	fmt.Println(model.DetailSelect()[0])
	c.HTML("Detail", core.J{
		"title":   "详情页",
		"details": model.DetailSelect(),
	})
}
