package handler

import "gateway/internal/web"

type HandlerMapping interface {
	GetHandler(exchange *web.ServerWebExchange) (Handler, error)
}
