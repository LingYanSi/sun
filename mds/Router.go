package mds

import (
	"fmt"
	"net/http"
	"sun/core"
	"time"

	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

type Router core.Router

// Routers 所有路由
type Routers map[string]Router

// Sun 请求处理
type SunRouter struct {
	routers         Routers
	notMatchJumpURL string
	redis           *redis.Client
}

// Add 添加路由
func (f *SunRouter) Add(key string, handler Router) {
	fmt.Println("添加路由: ", key)
	f.routers[key] = handler
}

// handleReq 处理请求
func (f *SunRouter) handleReq(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	// 输出请求信息
	fmt.Println(
		time.Now(),
		" |  request:",
		color.RedString(url),
		" method:",
		color.RedString(req.Method),
	)

	// 使用中间件机制，因为如果
	// 读取静态资源
	// 处理一般请求
	isMathed := false
	for key, handler := range f.routers {
		// 获取匹配到的url
		if isMath, _ := getRouter(key, url); isMath {
			isMathed = true
			fmt.Println("路由已匹配")
			handler.SetHTTP(req, res, f.redis)
			handler.Exec(handler, req.Method)
			break
		}
	}

	// 如果路由无匹配，则跳转到404页面
	if !isMathed {
		fmt.Println("路由未匹配，302跳转中")
		// 跳转到404页面
		res.Header().Set("Location", f.notMatchJumpURL)
		res.WriteHeader(302)
	}

}

// NotMatch 处理路由未匹配情况
func (f *SunRouter) NotMatch(path string) {
	f.notMatchJumpURL = path
}

func (f *SunRouter) init() {
	f.routers = Routers{} // 初始化slice，否则默认为nil
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	f.redis = client
}

// getRouter 获取路由
func getRouter(key, url string) (bool, map[string]string) {
	params := make(map[string]string)
	if key == url {
		return true, params
	}
	return false, params
}

// SRouter 需要获取Router
func SRouter() (core.Md, *SunRouter) {
	routers := &SunRouter{}
	routers.init()
	// 这里返回两个数据
	// Md 中间件
	// SunRouter 路由对象
	// 因为使用了闭包，可以使中间件获取到路由
	return core.MdCreate(func(req *http.Request, res http.ResponseWriter, next core.Next) core.Next {
		return func() {
			routers.handleReq(res, req)
			next()
		}
	}), routers
}
