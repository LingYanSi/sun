package controller

import (
	"sun/core"
)

// DeleteRedis 更新redis
type DeleteRedis struct {
	core.Base
}

// GET 处理get请求
func (dr *DeleteRedis) GET() {
	value := dr.Query("key")
	dr.Redis.Set("name", value, 0)
	// hash map 设置 key -> value 形式
	dr.Redis.HMSet("goods:sale", core.J{
		"name": "黑呵呵呵呵",
	})
	name := dr.Redis.HGet("goods:sale", "name").Val()
	// 允许跨域请求
	dr.Text("删除成功" + name)
}
