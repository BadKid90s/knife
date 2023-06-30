package web

type HandlerMapping interface {
	GetHandler(exchange *ServerWebExchange) (Handler, error)
}
