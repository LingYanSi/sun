package core

// 监听http请求，并根据

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"    // terminal字体颜色
	"github.com/go-redis/redis" // redis
)

// Routers 所有路由
type Routers map[string]Router

// Default 初始化
func Default() Sun {
	sun := Sun{}
	sun.init()
	return sun
}

// Sun 请求处理
type Sun struct {
	Routers
	NotMatchJumpURL string
	redis           *redis.Client
}

// Init 初始化
func (f *Sun) init() {
	f.Routers = make(map[string]Router)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	f.redis = client
}

// Add 添加路由
func (f *Sun) Add(key string, handler Router) {
	f.Routers[key] = handler
}

// handleReq 处理请求
func (f *Sun) handleReq(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	// 输出请求信息
	fmt.Println(
		time.Now(),
		" |  request:",
		color.RedString(url),
		" method:",
		color.RedString(req.Method),
	)

	isMathed := false
	for key, handler := range f.Routers {
		// 获取匹配到的url
		if isMath, _ := getRouter(key, url); isMath {
			isMathed = true
			handler.SetHTTP(req, res, f.redis)
			handler.Exec(handler, req.Method)
			break
		}
	}

	// 如果路由无匹配，则跳转到404页面
	if !isMathed {
		// 跳转到404页面
		res.Header().Set("Location", f.NotMatchJumpURL)
		res.WriteHeader(301)
	}
}

// NotMatch 处理路由未匹配情况
func (f *Sun) NotMatch(path string) {
	f.NotMatchJumpURL = path
}

// getRouter 获取路由
func getRouter(key, url string) (bool, map[string]string) {
	params := make(map[string]string)
	if key == url {
		return true, params
	}
	return false, params
}

// Listen 监听端都
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
