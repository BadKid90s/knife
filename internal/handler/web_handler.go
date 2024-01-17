package handler

import "gateway/internal/web"

// Handler 中间件处理器接口
type Handler interface {
	Handle(exchange *web.ServerWebExchange)
}

// HandlerFunc 中间件处理器接口的方法类型的实现
type HandlerFunc func(exchange *web.ServerWebExchange)

func (m HandlerFunc) Handle(exchange *web.ServerWebExchange) {
	m(exchange)
}
