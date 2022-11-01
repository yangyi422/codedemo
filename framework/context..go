package framework

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	// 是否超时标记
	hasTimeout bool
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// 返回Json
func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}
