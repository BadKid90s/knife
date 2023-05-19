package middleware

import "net/http"

type Handler interface {
	Handle(ctx *Context, writer http.ResponseWriter, request *http.Request) error
}

type HandlerFunc func(ctx *Context, writer http.ResponseWriter, request *http.Request) error

func (m HandlerFunc) Handle(ctx *Context, writer http.ResponseWriter, request *http.Request) error {
	return m(ctx, writer, request)
}
