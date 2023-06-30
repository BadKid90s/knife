package web

// Handler 中间件处理器接口
type Handler interface {
	Handle(exchange *ServerWebExchange)
}

// HandlerFunc 中间件处理器接口的方法类型的实现
type HandlerFunc func(exchange *ServerWebExchange)

func (m HandlerFunc) Handle(exchange *ServerWebExchange) {
	m(exchange)
}
