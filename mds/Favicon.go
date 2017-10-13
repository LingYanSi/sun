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
			reqPath := req.URL.Path
			if reqPath == "/favicon.ico" {
				file, err := os.Open("." + path)
				if err != nil {
					// 文件不存在
					res.Write([]byte("file not found"))
				} else {
					io.Copy(res, file)
				}
				defer file.Close()
			} else {
				next()
			}
		}
	})
}
