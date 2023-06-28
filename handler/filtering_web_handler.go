package handler

import "gateway/web"

func NewFilteringWebHandler() *FilteringWebHandler {
	return &FilteringWebHandler{}
}

type FilteringWebHandler struct {
}

func (h *FilteringWebHandler) Handle(exchange *web.ServerWebExchange) {

}
