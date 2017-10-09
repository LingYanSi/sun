package mds

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sun/core"
)

func getFileType(ext string) string {
	fileType := ""
	switch ext {
	case "css":
		fileType = "text/css"
	case "js":
		fileType = "application/javascript"
	case "html":
		fileType = "text/html"
	}
	return fileType
}

// Static 静态资源处理
func Static(dir ...string) core.Md {
	return core.MdCreate(func(req *http.Request, res http.ResponseWriter, next core.Next) core.Next {
		return func() {
			path := req.URL.Path
			match := false
			// 匹配文件夹
			for _, item := range dir {
				if strings.HasPrefix(path, item) {
					match = true
					break
				}
			}

			if match {
				fmt.Println("进入静态路由")
				// 打开文件
				file, err := os.Open("." + path)
				if err != nil {
					// 文件不存在
					res.Write([]byte("file not found"))
				} else {
					// 处理静态资源content-type
					ss := strings.Split(req.URL.Path, ".")
					ext := ss[len(ss)-1]
					if fileType := getFileType(ext); fileType != "" {
						res.Header().Set("content-type", fileType)
					}
					res.Header().Set("cache-control", "max-age=0")
					// 类似于node里的stream
					io.Copy(res, file)
				}
				// 关闭文件
				defer file.Close()
			} else {
				next()
			}
		}
	})
}
