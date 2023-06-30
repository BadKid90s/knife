package web

import (
	"net/http"
)

func NewServerWebExchange(write http.ResponseWriter, request *http.Request) *ServerWebExchange {
	return &ServerWebExchange{
		Write:      write,
		Request:    request,
		Attributes: map[string]interface{}{},
	}
}

type ServerWebExchange struct {
	Write      http.ResponseWriter
	Request    *http.Request
	Attributes map[string]interface{}
}
