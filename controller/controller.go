package controller

import (
	"context"
	"coredemo/framework"
	"time"
)

func ControllerHandler(ctx *framework.Context) error {
	timeoutCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	// 函数结束后调用cancel，结束整个context
	defer cancel()

	var finish = make(chan struct{}, 1)

	go func() {
		time.Sleep(10 * time.Second)
		ctx.Json(200, map[string]interface{}{
			"code": 0,
		})
		finish <- struct{}{}
	}()
	return ctx.Json(200, "ok")
}
