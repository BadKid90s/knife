package middleware

import "net/http"

type MiddlewareHandler interface {
	Handle(ctx *Context, writer http.ResponseWriter, request *http.Request) error
}

type MiddlewareHandlerFunc func(ctx *Context, writer http.ResponseWriter, request *http.Request) error
