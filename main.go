package main

import (
	"coredemo/framework"
	"net/http"
)

func main() {
	var server = &http.Server{
		// 自定义Handle
		Handler: framework.NewCore(),
		// 监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
}
