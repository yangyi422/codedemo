package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handle         ControllerHandler
	hasTimeout     bool        // 是否超时标记
	writerMux      *sync.Mutex // 写保护机制
}

// 创建一个新的context
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}
}

//=========================================
// base
// WriterMux
func (c *Context) WriterMux() {

}

// GetRequest
func (c *Context) GetRequest() {

}

// GetResponse
func (c *Context) GetResponse() {

}

// SetHasTimeout
func (c *Context) SetHasTimeout() {

}

// HasTimeout
func (c *Context) HasTimeout() bool {
	return c.hasTimeout
}

//=========================================
// context
// BaseContext
func (c *Context) BaseContext() context.Context {
	return c.request.Context()
}

func (c *Context) Deadline() {

}

func (c *Context) Done() <-chan struct{} {
	return c.BaseContext().Done()
}

func (c *Context) Err() {

}

func (c *Context) Value() {

}

//=========================================
// request
// QueryInt
func (c *Context) QueryInt() {

}

func (c *Context) QueryString() {

}

func (c *Context) QueryArray() {

}

func (c *Context) QueryAll() {

}

func (c *Context) FormInt() {

}

func (c *Context) FormString() {

}

func (c *Context) FormArray() {

}

func (c *Context) FormAll() {

}

func (c *Context) BindJson() {

}

//=========================================
// response
// Json 返回Json
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

func (c *Context) HTML() {

}

func (c *Context) Text() {

}
