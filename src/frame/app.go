// 控制器映射
// @by xiaoliang
// 仅供学习使用
package frame

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// 路由控制器 方法常量
const (
	defController = "index"
	defMethod     = "index"
)

// 路由表
type application struct {
	routes map[string]interface{}
}

// 初始化
func Default() *application {
	return &application{
		routes: make(map[string]interface{}),
	}
}

// 实现自定义路由
func (p *application) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 查找controller action
	controllerName, methodName := p.findControllerPathInfo(req)

	route, ok := p.routes[controllerName+"/"+strings.ToLower(methodName)]
	// 没找到控制器
	if !ok {
		http.NotFound(w, req)
		return
	}

	// 反射操作
	refV := reflect.ValueOf(route)
	if !refV.MethodByName(methodName).IsValid() {
		http.NotFound(w, req)
		return
		// fmt.Println("无效", methodName)
	}

	refV.Elem().FieldByName("Request").Set(reflect.ValueOf(req))
	refV.Elem().FieldByName("Response").Set(reflect.ValueOf(w))
	// 设置method
	refV.MethodByName(methodName).Call([]reflect.Value{})
}

// 查找控制器
// pathinfo
func (p *application) findControllerPathInfo(req *http.Request) (string, string) {
	path := req.URL.Path

	// 过滤后缀
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	pathinfo := strings.Split(path, "/")

	controllerName := defController
	// 控制器
	if len(pathinfo) > 1 {
		controllerName = pathinfo[1]
	}

	methodName := defMethod

	// 方法
	if len(pathinfo) > 1 {
		// 小写
		methodName = strings.Title(strings.ToLower(pathinfo[2]))
	}

	return controllerName, methodName
}

func (p *application) ln(str interface{}) {
	fmt.Printf("%v", str)
}

// Router 路由设置
func (p *application) GET(route string, controller interface{}) {
	p.routes[route] = controller
}

// run http
func (p *application) RUN(Addr string) error {
	fmt.Printf("%+v", p)
	// http.HandleFunc("/index", Index)
	fmt.Printf("listen on %s\n", Addr)
	return http.ListenAndServe(Addr, p)
}
