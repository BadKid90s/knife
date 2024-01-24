package knife

import (
	"math"
	"net/http"
)

const AbortIndex = math.MaxInt

var AllowTrueMiddlewareMatcher = MiddlewareMatcher(func(HttpResponse, HttpRequest) bool {
	return true
})

var (
	defaultHttpHandler http.Handler = http.NewServeMux()
)

type HttpHandleFunc func(http.ResponseWriter, *http.Request)

func (f HttpHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

type (
	HttpHandler struct {
		http.Handler
	}
	HttpResponse http.ResponseWriter
	HttpRequest  *http.Request

	MiddlewareMatcher func(HttpResponse, HttpRequest) bool

	MiddlewareFunc func(*Context)

	Context struct {
		Req    *http.Request
		Writer http.ResponseWriter
		index  int
		chain  *Chain
	}

	Chain struct {
		*HttpHandler
		middlewares        []MiddlewareFunc
		middlewareMatchers []MiddlewareMatcher
	}
)

func (c *Context) Next() {
	c.index++
	middlewares := c.chain.middlewares

	s := len(middlewares)
	for ; c.index < s; c.index++ {

		if c.chain.middlewareMatchers[c.index](c.Writer, c.Req) {
			middlewares[c.index](c)
		}
	}

	if c.index == s {
		c.chain.HttpHandler.ServeHTTP(c.Writer, c.Req)
	}

}
func (c *Context) Abort(code int) {
	c.Writer.WriteHeader(code)
	c.index = AbortIndex
}

func (c *Context) Fail(code int, err error) {
	c.Abort(code)
	_, _ = c.Writer.Write([]byte(err.Error()))
}

func NewChain(httpHandle http.Handler, handlers ...MiddlewareFunc) *Chain {
	var h *HttpHandler
	if httpHandle == nil {
		h = &HttpHandler{defaultHttpHandler}
	} else {
		h = &HttpHandler{httpHandle}
	}
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
		Writer: w,
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
