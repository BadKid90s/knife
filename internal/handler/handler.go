package handler

import "gateway/internal/web"

// Handler 中间件处理器接口
type Handler interface {
	Handle(exchange *web.ServerWebExchange)
}

// Func 中间件处理器接口的方法类型的实现
type Func func(exchange *web.ServerWebExchange)

func (m Func) Handle(exchange *web.ServerWebExchange) {
	m(exchange)
}
