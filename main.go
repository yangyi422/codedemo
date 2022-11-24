package main

import (
	"coredemo/framework"
	"coredemo/router"
	"net/http"
)

func main() {
	var core = framework.NewCore()
	router.RegisterRouter(core)
	var server = &http.Server{
		// 自定义Handle
		Handler: core,
		// 监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}
