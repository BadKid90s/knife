package knife

import (
	"math"
	"net/http"
)

const AbortIndex = math.MaxInt

type Context struct {
	Req    *http.Request
	Writer HttpResponseWriter
	index  int
	chain  *Chain
}

func (c *Context) Next() {
	c.index++
	middlewares := c.chain.middlewares

	s := len(middlewares)
	for ; c.index < s; c.index++ {

		if c.chain.middlewareMatchers[c.index](c.Writer, c.Req) {
			middlewares[c.index](c)
		}
	}

	if c.index == s && c.chain.HttpHandler.Handler != nil {
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
