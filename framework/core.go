package framework

import (
	"log"
	"net/http"
	"strings"
)

// 动态路由
type Core struct {
	router map[string]*Tree // 一级map
}

func NewCore() *Core {
	// 初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error:", err)
	}
}

func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler {
	// uri和method转化为大写
	uri := request.URL.Path
	method := request.Method
	uri = strings.ToUpper(uri)
	method = strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[method]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// // 框架核心结构实现Handle接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := NewContext(request, response)
	router := c.FindRouterByRequest(request)
	// 如果没有找到路由对应的handle，返回404错误
	if router == nil {
		ctx.Json(404, "not found")
		return
	}

	// 如果调用的函数返回err，表示存在程序内部错误，返回500错误
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}

//======================================
// 静态化路由
// 框架核心结构
// type Core struct {
// 	router map[string]map[string]ControllerHandler // 一级map
// }

// // 初始化框架核心结构
// func NewCore() *Core {
// 	// 定义二级map
// 	getRouter := map[string]ControllerHandler{}
// 	postRouter := map[string]ControllerHandler{}
// 	putRouter := map[string]ControllerHandler{}
// 	deleteRouter := map[string]ControllerHandler{}

// 	// 将二级map写入一级map
// 	router := map[string]map[string]ControllerHandler{}
// 	router["GET"] = getRouter
// 	router["POST"] = postRouter
// 	router["PUT"] = putRouter
// 	router["DELETE"] = deleteRouter
// 	return &Core{router}
// }

// // 框架核心结构实现Handle接口
// func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
// 	ctx := NewContext(request, response)
// 	router := c.FindRouterByRequest(request)
// 	// 如果没有找到路由对应的handle，返回404错误
// 	if router == nil {
// 		ctx.Json(404, "not found")
// 		return
// 	}

// 	// 如果调用的函数返回err，表示存在程序内部错误，返回500错误
// 	if err := router(ctx); err != nil {
// 		ctx.Json(500, "inner error")
// 		return
// 	}
// }

// //=========================================
// // 注册路由
// // Get 对应get请求
// func (c *Core) Get(url string, handler ControllerHandler) {
// 	upperUrl := strings.ToUpper(url)
// 	c.router["GET"][upperUrl] = handler
// }

// // Post 对应post请求
// func (c *Core) Post(url string, handler ControllerHandler) {
// 	upperUrl := strings.ToUpper(url)
// 	c.router["POST"][upperUrl] = handler
// }

// // Put 对应put请求
// func (c *Core) Put(url string, handler ControllerHandler) {
// 	upperUrl := strings.ToUpper(url)
// 	c.router["PUT"][upperUrl] = handler
// }

// // Delete 对应delete请求
// func (c *Core) Delete(url string, handler ControllerHandler) {
// 	upperUrl := strings.ToUpper(url)
// 	c.router["DELETE"][upperUrl] = handler
// }

// //=========================================
// // 匹配路由
// func (c *Core) FindRouterByRequest(request *http.Request) ControllerHandler {
// 	// 转化大写，保证大小写不敏感
// 	uri := request.URL.Path
// 	method := request.Method
// 	upperUri := strings.ToUpper(uri)
// 	upperMethod := strings.ToUpper(method)

// 	// 遍历路由map，返回handle
// 	if methodHandles, ok := c.router[upperMethod]; ok {
// 		if handle, ok := methodHandles[upperUri]; ok {
// 			return handle
// 		}
// 	}
// 	return nil
// }
