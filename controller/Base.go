package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"sun/util"
)

// 正则匹配等规则 -> 对应一个struct struct实现get/post/all 等方法

// Router 基类
// 接口用来判断对象是否实现了其声明的所有方法
// 任何数据都实现了一个空接口 interface{}
type Router interface {
	JSON(j J)
	HTML(path string, data interface{})
	Text(content string)
	SetHeader(key, value string)
	GET()
	POST()
	SetHTTP(req *http.Request, res http.ResponseWriter)
	Exec(r Router, method string)
}

// J map类型
type J map[string]interface{}

// Base handle haha
// struct相当于class 实例化也很简单 Base{Name: "bulalla"}
type Base struct {
	Req     *http.Request
	Res     http.ResponseWriter
	Cookies http.Cookie
	// Name string // 注意以大写字母开头的属性才能被外部访问，类似于public Name
	// age  int    // 小写字母的只能被内部访问 类似于private age
}

// SetHTTP 设置上下文环境
func (t *Base) SetHTTP(req *http.Request, res http.ResponseWriter) {
	t.Req = req
	t.Res = res
}

// Exec 处理请求
func (t *Base) Exec(c Router, method string) {
	switch method {
	case "GET":
		c.GET()
	case "POST":
		c.POST()
	}
}

// GET 返回html
func (t *Base) GET() {
	// 加载模板然后渲染
	t.Text("success")
}

// POST 返回html
func (t *Base) POST() {
	// 加载模板然后渲染
	t.Text("success")
}

// HTML 返回html
func (t *Base) HTML(path string, data interface{}) {
	// 加载模板然后渲染
}

// JSON 返回json
func (t *Base) JSON(data J) {
	// 加载模板然后渲染
	t.Res.WriteHeader(200)
	// 为什么这里有需要声明指针？
	j, err := json.Marshal(data)
	if err == nil {
		t.Res.Write(j)
	} else {
		fmt.Println(err)
	}
}

// Text return
func (t *Base) Text(content string) {
	// fmt.Println(
	// 	t.Req.Method,
	// 	t.Req.URL,
	// 	t.Req.UserAgent(),
	// 	t.Req.Header,
	// )

	// 声明code
	t.Res.WriteHeader(200)
	// 为什么这里有需要声明指针？
	t.Res.Write([]byte(content))
}

// SetHeader 添加header
func (t *Base) SetHeader(key string, value string) {
	header := t.Res.Header()
	header.Set(key, value)
}

// GetCookie 回去cookie
func (t *Base) GetCookie(key string) string {
	c, err := t.Req.Cookie(key)
	if err == nil {
		return c.Value
	}
	return ""
}

// SetCookie 设置cookie
func (t *Base) SetCookie(key string, value string, maxAge int) {
	cookie := http.Cookie{
		Name:  key,
		Value: value,
		Path:  "/",
		// Domain:   "",
		MaxAge:   maxAge,
		Secure:   false,
		HttpOnly: false,
		Raw:      "",
		// Unparsed []string // Raw text of unparsed attribute-value pairs
	}
	http.SetCookie(t.Res, &cookie)
}

// DeleteCookie 删除指定cookie
func (t *Base) DeleteCookie(key string) {
	t.SetCookie(key, "", -1)
}

// ClearCookie 清空cookie
func (t *Base) ClearCookie() {
	for _, c := range t.Req.Cookies() {
		t.DeleteCookie(c.Name)
	}
}

// allow 允许跨域请求
func (c *Base) allow() {
	referer := c.Req.Referer()
	handler := getHander(referer)
	hosts := []string{"127.0.0.1", "[^\\.]*\\.weipaitang\\.com"}
	matched := arrSome(hosts, handler)
	if matched {
		c.SetHeader("Access-Control-Allow-Credentials", "true")
		c.SetHeader("Access-Control-Allow-Headers", "Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")
		c.SetHeader("Access-Control-Allow-Methods", "GET,HEAD,PUT,POST,DELETE")
		urlResult, err := url.Parse(referer)
		if err == nil {
			port := urlResult.Port()
			port = util.Three(port != "", ":", "").(string) + port
			c.SetHeader("Access-Control-Allow-Origin", urlResult.Scheme+"://"+urlResult.Hostname()+port)
		} else {
			c.SetHeader("Access-Control-Allow-Origin", "*")
		}
	}
}

// ArrSomeHander 处理数组
type ArrSomeHander func(item string, index int) bool

// arrSome 只要有一个符合条件
func arrSome(arr []string, m ArrSomeHander) bool {
	matched := false
	for index, value := range arr {
		matched = m(value, index)
		if matched {
			break
		}
	}
	return matched
}

// arrEvery 都符合条件
func arrEvery(arr []string, m ArrSomeHander) bool {
	matched := false
	for index, value := range arr {
		matched = m(value, index)
		if !matched {
			break
		}
	}
	return matched
}

// arrFilter 都符合条件
func arrFilter(arr []string, m ArrSomeHander) []string {
	var newArr []string
	matched := false
	for index, value := range arr {
		matched = m(value, index)
		if matched {
			newArr[len(newArr)] = value
		}
	}
	return newArr
}

func getHander(referer string) ArrSomeHander {
	return func(item string, index int) bool {
		return regexp.MustCompile("^https?://"+item).FindString(referer) != ""
	}
}

// Routers 所有路由
type Routers map[string]Router

// Sun 请求处理
type Sun struct {
	Routers
}

// Init 初始化
func (f *Sun) Init() {
	f.Routers = make(map[string]Router)
}

// AddRouter 添加路由
func (f *Sun) AddRouter(key string, handler Router) {
	f.Routers[key] = handler
}

// HandleReq 处理请求
func (f *Sun) HandleReq(req *http.Request, res http.ResponseWriter) {
	url := req.URL.Path
	for key, handler := range f.Routers {
		if key == url {
			handler.SetHTTP(req, res)
			handler.Exec(handler, req.Method)
			break
		}
	}
}
