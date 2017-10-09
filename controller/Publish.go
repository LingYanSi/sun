package controller

import (
	"encoding/json"
	"fmt"
	"sun/core"
)

type M struct {
	Name  string `json:"name"` // 需求名称
	Pics  string `json:"pics"` // 图片
	Des   string `json:"des"`  // 描述
	Tel   string // 电话
	Price int    // 价格
	Time  string // 周期
	Doc   string // 文档
	Ui    string // ui
	City  string // 城市
}

// Index 处理主页
type Publish struct {
	core.Base
}

// GET 处理get请求
func (c *Publish) GET() {
	c.HTML("publish", core.J{
		"title": "需求发布",
	})
}

func (c *Publish) handleErr(msg string) {
	c.JSON(core.J{
		"code": "400",
		"msg":  msg,
	})
}

func getString(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	default:
		return ""
	}
}

func (this *Publish) POST() {
	model := M{}
	// 数据校验
	if model.Name = getString(this.Input("name")); model.Name == "" {
		this.handleErr("项目标题不能为空")
		return
	}

	if model.Pics = getString(this.Input("pics")); model.Pics == "" {
		this.handleErr("项目图片不能为空")
		return
	}

	if model.Des = getString(this.Input("des")); model.Des == "" {
		this.handleErr("项目描述不能为空")
		return
	}

	result, err := json.Marshal(model)
	if err != nil {
		fmt.Println("json encode失败", err)
	}
	fmt.Println("获取result: ", string(result))
	this.Redis.LPush("lists", string(result))

	this.Redirect("/list")
	// this.JSON(core.J{
	// 	"code": 200,
	// 	"msg":  "成功",
	// })
}
