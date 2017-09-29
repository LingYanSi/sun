package controller

import (
	"sun/core"
)

// NotFound 找不到服务
type NotFound struct {
	core.Base
}

// Any 返回404
func (t *NotFound) Any() {
	t.HTML("404", core.J{
		"title": "找不到",
		"body":  "拔剑四顾心茫然",
	})
}
