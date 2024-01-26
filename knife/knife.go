package knife

import (
	"net/http"
)

type (
	HttpHandler struct {
		http.Handler
	}

	Chain struct {
		*HttpHandler
		middlewares        []MiddlewareFunc
		middlewareMatchers []MiddlewareMatcher
	}
)

func NewChain(httpHandle http.Handler, handlers ...MiddlewareFunc) *Chain {
	var h = &HttpHandler{httpHandle}
	chain := newChain(h, nil, nil)
	return chain.Use(handlers...)
}

func newChain(httpHandle *HttpHandler, middlewares []MiddlewareFunc, middlewareMatchers []MiddlewareMatcher) *Chain {
	return &Chain{
		HttpHandler:        httpHandle,
		middlewares:        append([]MiddlewareFunc{}, middlewares...),
		middlewareMatchers: append([]MiddlewareMatcher{}, middlewareMatchers...),
	}
}

func (chain *Chain) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	chain.serverHandle().ServeHTTP(w, req)
}

func (chain *Chain) createContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: NewResponseWriter(w),
		Req:    req,
		index:  -1,
		chain:  chain,
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
	var newHandlers []MiddlewareFunc
	newHandlers = append(newHandlers, chain.middlewares...)
	newHandlers = append(newHandlers, middlewares...)

	var middlewareMatchers []MiddlewareMatcher
	middlewareMatchers = append(middlewareMatchers, chain.middlewareMatchers...)
	for i := 0; i < len(middlewares); i++ {
		middlewareMatchers = append(middlewareMatchers, matcher)
	}
	chain.middlewares = newHandlers
	chain.middlewareMatchers = middlewareMatchers
	return chain
}
