package knife

import (
	"net/http"
)

type (
	Chain struct {
		middlewares []*Middleware
	}
)

func NewChain(handlers ...MiddlewareFunc) *Chain {
	chain := &Chain{middlewares: make([]*Middleware, 0)}
	return chain.Use(handlers...)
}

func (chain *Chain) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	chain.serverHandle().ServeHTTP(w, req)
}

func (chain *Chain) createContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:      NewResponseWriter(w),
		Req:         req,
		index:       -1,
		middlewares: chain.middlewares,
	}
}

func (chain *Chain) serverHandle() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		context := chain.createContext(writer, request)
		context.Next()
	})
}

func (chain *Chain) Use(middlewares ...MiddlewareFunc) *Chain {

	return chain.UseMatcher(AllowTrueMiddlewareMatcher, middlewares...)
}

func (chain *Chain) UseMatcher(matcher MiddlewareMatcher, middlewares ...MiddlewareFunc) *Chain {
	for _, middlewareFunc := range middlewares {
		middleware := NewMiddleware(matcher, middlewareFunc)
		chain.middlewares = append(chain.middlewares, middleware)
	}
	return chain
}
