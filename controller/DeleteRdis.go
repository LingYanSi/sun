package controller

import (
	"sun/core"
	"sun/model"
)

// DeleteRedis 更新redis
type DeleteRedis struct {
	core.Base
}

// GET 处理get请求
func (dr *DeleteRedis) GET() {
	value := dr.Query("key")
	model.Redis.Set("name", value, 0)
	// hash map 设置 key -> value 形式
	model.Redis.HMSet("goods:sale", core.J{
		"name": "黑呵呵呵呵",
	})
	name := model.Redis.HGet("goods:sale", "name").Val()
	// 允许跨域请求
	dr.Text("删除成功" + name)
}
