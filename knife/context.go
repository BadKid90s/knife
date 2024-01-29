package knife

import (
	"math"
	"net/http"
)

const AbortIndex = math.MaxInt / 2

type Context struct {
	Req         *http.Request
	Writer      HttpResponseWriter
	index       int
	middlewares []*Middleware
}

func (c *Context) Next() {
	c.index++
	middlewares := c.middlewares

	s := len(middlewares)
	for ; c.index < s; c.index++ {
		middleware := c.middlewares[c.index]
		if middleware.matcher(c.Writer, c.Req) {
			middleware.handler(c)
		}
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
