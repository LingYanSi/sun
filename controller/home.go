package controller

// Home 处理主页
type Home struct {
	Base
}

type goods struct {
	URI string `json:"-"`
	Age int    `json:"age,omitempty"`
}

// GET 处理get请求
func (c *Home) GET() {
	// 允许跨域请求
	c.allow()

	c.SetHeader("add", "bbb")
	c.SetHeader("add", "---")
	// c.Text("哈哈哈哈哈哈哈")
	c.JSON(J{
		"Name": 1,
		"age":  []int{1, 2, 3, 4},
		"goods": goods{
			URI: "weeee",
		},
		"arr": []int{},
	})
}
