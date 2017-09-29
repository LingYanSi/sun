package controller

import (
	"fmt"
	"strings"
	"sun/core"
	"sun/util"
)

// Home 处理主页
type Comments struct {
	core.Base
}

// POST 处理get请求
func (this *Comments) POST() {
	content := this.Input("content")
	fmt.Println("获取输入", content, this.Req.Header.Get("Content-Type"))

	// xx.(type)用来判断数据类型，并返回对应的值
	// 使用switch来校验类型，在Go中switch默认break，如果不想使用break可以使用fallThrough关键字
	switch c := content.(type) {
	case string:
		// 去除空格
		if strings.TrimSpace(c) != "" {
			this.Redis.LPush("comments", c)
		} else {
			fmt.Println("content is empty!")
		}
	}

	// 允许跨域请求
	this.JSON(core.J{
		"code": 0,
	})
}

// GET 处理get请求
func (this *Comments) GET() {
	comments := this.Redis.LRange("comments", 0, 100).Val()
	// 反转数组
	comments = util.Reverse(comments)

	this.JSON(core.J{
		"list": comments,
	})
}

/*
* api 处理的话，就是返回json
* 如何处理
 */
