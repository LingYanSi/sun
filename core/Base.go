package core

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sun/util"
	"time"
)

// 正则匹配等规则 -> 对应一个struct struct实现get/post/all 等方法

// J map类型
type J map[string]interface{}

// Router 基类
// 接口用来判断对象是否实现了其声明的所有方法
// 任何数据都实现了一个空接口 interface{}
type Router interface {
	Init() // 初始化
	JSON(j J)
	HTML(path string, data J)
	Text(content string)
	GET()
	POST()
	Any()
	SetHeader(key, value string)
	SetHTTP(req *http.Request, res http.ResponseWriter)
	Exec(r Router, method string)
}

// Input 获取数据
type Input func(string) interface{}

// Base handle haha
// struct相当于class 实例化也很简单 Base{Name: "bulalla"}
type Base struct {
	Req        *http.Request
	Res        http.ResponseWriter
	Cookies    http.Cookie
	IP         string // 注意以大写字母开头的属性才能被外部访问，类似于public Name
	URL        string // 请求path
	Protocol   string // 协议
	Host       string // host
	Path       string // 请求路径
	RequestURI string // 带search的path
	Input      Input  // 获取客户端传递的数据
	isNotAny   bool   // 小写字母的只能被内部访问 类似于private age
}

// SetHTTP 设置上下文环境
func (c *Base) SetHTTP(req *http.Request, res http.ResponseWriter) {
	c.Req = req
	c.Res = res
}

// Exec 处理请求
func (c *Base) Exec(router Router, method string) {
	c.Protocol = util.Three(strings.Contains(c.Req.Proto, "1."), "http", "https").(string) // 协议
	c.Host = c.Req.Host                                                                    // host 包含port
	c.Path = c.Req.URL.Path                                                                // 不含 ?11=
	c.RequestURI = c.Req.RequestURI                                                        // 含 ?11=
	c.URL = c.Protocol + "://" + c.Host + c.Req.RequestURI
	c.IP = c.Req.RemoteAddr // 客户端ip地址

	c.Input = c.getPostData() // 获取请求数据

	// 返回status code，需要注意的是code需要在header.set之后射入
	// c.Res.WriteHeader(200)
	// 缓存控制
	// c.SetHeader("Cache-Control", "max-age=0")
	// 初始化请求，可用于log等等等
	router.Init()
	// 有过有Any实现就使用，反则调用相应的Method实现
	if router.Any(); c.isNotAny {
		switch method {
		case "GET":
			router.GET()
		case "POST":
			router.POST()
		}
	}
}

// Init 执行初始化
func (c *Base) Init() {

}

// GET 返回html
func (c *Base) GET() {
	// 加载模板然后渲染
	c.Res.WriteHeader(404)
	// c.JSON(J{
	// 	"code": "404",
	// 	"msg":  "找不到服务",
	// })
}

// POST 返回html
func (c *Base) POST() {
	// 加载模板然后渲染
	c.Res.WriteHeader(404)
}

// DELETE 删除数据
func (c *Base) DELETE() {
	c.Res.WriteHeader(404)
}

// PUT 更新数据
func (c *Base) PUT() {
	c.Res.WriteHeader(404)
}

// Any 对对对
func (c *Base) Any() {
	// 加载模板然后渲染
	c.isNotAny = true
}

// Redirect 处理跳转
func (c *Base) Redirect(path string) {
	c.SetHeader("Location", path)
	c.Res.WriteHeader(302)
}

// HTML 返回html
func (c *Base) HTML(path string, data J) {
	// 加载模板然后渲染
	c.SetHeader("Content-Type", "text/html; charset=utf8")
	filename := "views/" + path + ".html"
	header := "views/common/header.html"
	footer := "views/common/footer.html"
	// 添加函数，用来执行
	tmpl := template.New("index").Funcs(template.FuncMap{
		"getPath": func(path string) string {
			fmt.Println("时间戳类型: ", reflect.TypeOf(time.Now().Unix()))
			return path + "?version=" + strconv.FormatInt(time.Now().Unix(), 10)
		},
	})
	tmpl, err := tmpl.ParseFiles(header, footer, filename)
	// Error checking elided
	if err != nil {
		fmt.Println("模板解析出错: "+filename, err)
	}
	// 如果使用PareseFiles解析文件模板，需要使用ExecuteTemplate来进行数据渲染
	err = tmpl.ExecuteTemplate(
		c.Res,
		path+".html",
		data,
	)
	if err != nil {
		fmt.Println("模板渲染出错: "+filename, err)
	}
}

// JSON 返回json
func (c *Base) JSON(data J) {
	// 加载模板然后渲染
	c.SetHeader("Content-Type", "text/html; charset=utf8")
	// 为什么这里有需要声明指针？
	j, err := json.Marshal(data)
	if err == nil {
		c.Res.Write(j)
	} else {
		fmt.Println("json整理出错: ", err)
	}
}

// Text return
func (c *Base) Text(content string) {
	// 声明code
	c.SetHeader("Content-Type", "text/plain; charset=utf8")
	// 为什么这里有需要声明指针？
	c.Res.Write([]byte(content))
}

func mergeMap(map1 J, map2 J) interface{} {
	newMap := J{}
	for key, value := range map1 {
		newMap[key] = value
	}
	for key, value := range map2 {
		newMap[key] = value
	}
	return newMap
}

// getPostData 获取post请求数据
func (c *Base) getPostData() Input {
	// 解析请求参数
	c.Req.ParseForm()
	var t J
	req := c.Req
	// 如果Content-Type 包含json字段，
	if strings.Contains(req.Header.Get("Content-Type"), "text/json") || strings.Contains(req.Header.Get("content-type"), "text/json") {
		// 解析json，对于post数据的获取需要先ParseForm
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&t)
		if err != nil {
			// panic(err)
		}
		defer req.Body.Close()
	}

	return func(key string) interface{} {
		value := t[key]
		if value == nil {
			value = c.Query(key)
		}
		return value
	}
}

// Query 获取form/query
func (c *Base) Query(key string) string {
	// fmt.Println("post: ", getPostData(c.Req))
	return c.Req.Form.Get(key)
}

// SetHeader 添加header
func (c *Base) SetHeader(key string, value string) {
	header := c.Res.Header()
	header.Set(key, value)
}

// GetCookie 回去cookie
func (c *Base) GetCookie(key string) string {
	cookie, err := c.Req.Cookie(key)
	if err == nil {
		return cookie.Value
	}
	return ""
}

// SetCookie 设置cookie
func (c *Base) SetCookie(key string, value string, maxAge int) {
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
	http.SetCookie(c.Res, &cookie)
}

// DeleteCookie 删除指定cookie
func (c *Base) DeleteCookie(key string) {
	c.SetCookie(key, "", -1)
}

// ClearCookie 清空cookie
func (c *Base) ClearCookie() {
	for _, cookie := range c.Req.Cookies() {
		c.DeleteCookie(cookie.Name)
	}
}

// Allow 允许跨域请求
func (c *Base) Allow() {
	referer := c.Req.Referer()
	hosts := []string{"127.0.0.1", "[^\\.]*\\.weipaitang\\.com"}

	matched := util.ArrSome(hosts, func(item string, index int) bool {
		return regexp.MustCompile("^https?://"+item).FindString(referer) != ""
	})

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
