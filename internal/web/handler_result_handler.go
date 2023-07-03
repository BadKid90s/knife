package web

type HandlerResultHandler interface {
	handleResult(exchange *ServerWebExchange, result *HandlerResult)
}
