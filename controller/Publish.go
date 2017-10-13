package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sun/core"
	"sun/model"
)

type M struct {
	Name  string `json:"name"`  // 需求名称
	Pics  string `json:"pics"`  // 图片
	Des   string `json:"des"`   // 描述
	Tel   string `json:"tel"`   // 电话
	Price int    `json:"price"` // 价格
	Time  string `json:"time"`  // 周期
	Doc   string `json:"doc"`   // 文档
	UI    string `json:"ui"`    // ui
	City  string `json:"city"`  // 城市
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

func getInt(data interface{}) int {
	switch data.(type) {
	case int:
		return data.(int)
	case string:
		num, err := strconv.Atoi(data.(string))
		if err != nil {
			return -1
		}
		return num
	default:
		return -1
	}
}

func (this *Publish) POST() {
	data := M{}
	// 数据校验
	if data.Name = getString(this.Input("name")); data.Name == "" {
		this.handleErr("项目标题不能为空")
		return
	}

	if data.Pics = getString(this.Input("pics")); data.Pics == "" {
		this.handleErr("项目图片不能为空")
		return
	}

	if data.Des = getString(this.Input("des")); data.Des == "" {
		this.handleErr("项目描述不能为空")
		return
	}

	if data.Tel = getString(this.Input("tel")); data.Tel == "" {
		this.handleErr("手机号码不能为空")
		return
	}

	if data.Price = getInt(this.Input("price")); data.Price == -1 {
		this.handleErr("价格不能为空")
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json encode失败", err)
	}
	fmt.Println("获取result: ", string(result))
	model.Redis.LPush("lists", string(result))

	this.Redirect("/list")
	// this.JSON(core.J{
	// 	"code": 200,
	// 	"msg":  "成功",
	// })
}
