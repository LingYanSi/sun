package controller

import (
	"sun/core"
	"sun/model"
)

// Home 处理主页
type SJSON struct {
	core.Base
}

// Any 处理get请求
func (this *SJSON) POST() {
	model.Redis.Set("name", "个屁", 0)
	// 允许跨域请求
	this.JSON(core.J{
		"Name": 1,
		"age":  []int{1, 2, 3, 4},
		"goods": goods{
			URI: "we eee",
		},
		"arr": []int{},
		"etc": "微拍堂11111",
	})
}

/*
* api 处理的话，就是返回json
* 如何处理
 */
