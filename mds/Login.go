package mds

import (
	"net/http"
	"strings"
	"sun/core"
)

// Login 处理登录接口
func Login() core.Md {
	return core.MdCreate(func(req *http.Request, res http.ResponseWriter, next core.Next) core.Next {
		return func() {
			path := req.URL.Path
			if strings.HasSuffix(path, "-l") {
				// 判断是否含有指定cookie
				res.Write([]byte("请求登录接口"))
				// 判断cookie是否正确 不振孤噩就删除cookie
			} else {
				next()
			}
		}
	})
}
