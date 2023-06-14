package middleware

import "net/http"

// Handler 中间件处理器接口
type Handler interface {
	Handle(ctx *Context, writer http.ResponseWriter, request *http.Request) (bool, error)
}

// HandlerFunc 中间件处理器接口的方法类型的实现
type HandlerFunc func(ctx *Context, writer http.ResponseWriter, request *http.Request) (bool, error)

func (m HandlerFunc) Handle(ctx *Context, writer http.ResponseWriter, request *http.Request) (bool, error) {
	return m(ctx, writer, request)
}
