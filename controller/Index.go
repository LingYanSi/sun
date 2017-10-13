package controller

import (
	"fmt"
	"sun/core"
	"sun/model"
)

// Index 处理主页
type Index struct {
	core.Base
}

type goods struct {
	URI string `json:"-"` // - 表示忽略不输出
	// struct内大写开头表示公共属性可被访问，小写为私有属性不可被外部访问
	// 我们希望输出的时候显示小写，omitempty表示数据为空时不显示
	Age int `json:"age,omitempty"`
}

func (c *Index) getData() int64 {
	return 100
}

// Init
func (c *Index) Init() {
	fmt.Println("初始化调用-------")
}

// GET 处理get请求
func (c *Index) GET() {
	// 允许跨域请求
	c.Allow()
	// c.SetHeader("Content-Type", "text/html; charset=utf8")
	// c.Res.WriteHeader(200)

	name, err := model.Redis.Get("name").Result()
	if err != nil {
		name = "出错了"
	}
	fmt.Println("redis name:", name)
	c.HTML("index", core.J{
		"title": "首页",
		"body":  "我是body: " + name,
		"data": core.J{
			"title": "哈哈哈哈哈",
		},
	})
}

// POST 处理post请求
func (c *Index) POST() {
	c.JSON(core.J{
		"card": 10000,
	})
}

// func (c *Index) Any() {
// 	// 数值转字符串
// 	s := strconv.FormatInt(c.getData(), 10)
// 	fmt.Println("input: ", c.Input("mm"), c.URL, c.Path)
// 	// c.Redirect("https://www.baidu.com")

// 	// chan 用于协程通信 channel <- 向协程传递数据
// 	// a := <- channel 接收数据
// 	// 可用于数据查询，网络请求等
// 	channel := make(chan string)
// 	// go 关键字表示，开启一个协程，处理异步任务
// 	go func(channel chan string) {
// 		time.AfterFunc(time.Second*1, func() {
// 			fmt.Println("的是发送到发送到发生的")
// 			channel <- "我爱你呀"
// 		})
// 	}(channel)
// 	str := <-channel

// 	fmt.Println(str)
// 	c.Text("你打的优雅" + s + c.Query("name") + " ip:" + c.IP)
// }
