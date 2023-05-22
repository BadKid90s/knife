package middleware

import "net/http"

// Handler 中间价接口
type Handler interface {
	Handle(ctx *Context, writer http.ResponseWriter, request *http.Request) error
}

// HandlerFunc 中间价接口方法的实现
type HandlerFunc func(ctx *Context, writer http.ResponseWriter, request *http.Request) error

func (m HandlerFunc) Handle(ctx *Context, writer http.ResponseWriter, request *http.Request) error {
	return m(ctx, writer, request)
}
