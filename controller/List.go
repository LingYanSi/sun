package controller

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sun/core"
)

// Index 处理主页
type List struct {
	core.Base
}

// GET 处理get请求
func (this *List) GET() {
	list := this.Redis.LRange("lists", 0, 100).Val()
	var newList []interface{}
	for _, item := range list {
		i := core.J{}
		err := json.Unmarshal([]byte(item), &i)
		newList = append(newList, i)
		fmt.Println("列表", err, reflect.TypeOf(i))
	}

	this.HTML("list", core.J{
		"title": "所有需求",
		"list":  newList,
	})
}
