package mds

import (
	"io"
	"net/http"
	"os"
	"sun/core"
)

// Login 处理登录接口
func Favicon(path string) core.Md {
	return core.MdCreate(func(req *http.Request, res http.ResponseWriter, next core.Next) core.Next {
		return func() {
			path := req.URL.Path
			if path == "/favicon.ico" {
				file, err := os.Open("." + path)
				if err != nil {
					next()
				} else {
					io.Copy(res, file)
				}
			} else {
				next()
			}
		}
	})
}
