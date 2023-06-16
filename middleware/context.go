package middleware

import (
	"log"
	"net/http"
)

type Context struct {
	//有序的中间件
	handlers []Handler
	//中间件执行索引
	index int
}

func (c *Context) Next(write http.ResponseWriter, request *http.Request) {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		err := c.handlers[c.index].Handle(c, write, request)
		if err != nil {
			log.Println("middleware running error")
		}
	}
}

func createContext(handlers []Handler) *Context {
	return &Context{
		index:    -1,
		handlers: handlers,
	}
}
