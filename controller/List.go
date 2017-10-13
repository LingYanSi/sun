package controller

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sun/core"
	"sun/model"
)

// Index 处理主页
type List struct {
	core.Base
}

// GET 处理get请求
func (this *List) GET() {
	list := model.Redis.LRange("lists", 0, 100).Val()
	var newList []interface{}
	for _, item := range list {
		i := core.J{}
		err := json.Unmarshal([]byte(item), &i)
		newList = append(newList, i)
		fmt.Println("列表", err, reflect.TypeOf(i))
	}

	fmt.Println(newList)

	this.HTML("list", core.J{
		"title": "所有需求",
		"list":  newList,
	})
}
