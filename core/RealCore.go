package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type Sun struct {
	mds []Md
	Req *http.Request
	Res http.ResponseWriter
	Ctx map[string]interface{}
}

// Next 函数
type Next func()

// addCtx 添加上下文
type addCtx func(req *http.Request, res http.ResponseWriter) Next

// Md 中间件
type Md func(n Next) addCtx

// Handler 创建一个中间件
type Handler func(req *http.Request, res http.ResponseWriter, next Next) Next

// MdCreate 创建中间件
func MdCreate(handler Handler) Md {
	return func(next Next) addCtx {
		// 接受其他公共参数
		return func(req *http.Request, res http.ResponseWriter) Next {
			// 执行实际函数
			return handler(req, res, next)
		}
	}
}

// Use 添加中间件
func (m *Sun) Use(mds ...Md) *Sun {
	if len(m.mds) == 0 {
		var initMds []Md
		m.mds = initMds
	}
	m.mds = append(m.mds, mds...)
	return m
}

func (m *Sun) md() {}

func (m *Sun) combine() Next {
	md := m.md
	length := len(m.mds) - 1
	for index := range m.mds {
		md = m.mds[length-index](md)(m.Req, m.Res)
	}
	return md
}

func (m *Sun) handleReq(res http.ResponseWriter, req *http.Request) {
	m.Req = req
	m.Res = res
	m.combine()()
}

func (f *Sun) Listen(port string) {
	// log.Fatal("sun run: http://localhost:" + port)
	fmt.Println("sun run:", color.GreenString("http://localhost:"+port))
	// 添加http处理回调函数
	http.HandleFunc("/", f.handleReq)
	// 监听端口
	err := http.ListenAndServe(":"+port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Default() *Sun {
	sun := Sun{}
	return &sun
}
