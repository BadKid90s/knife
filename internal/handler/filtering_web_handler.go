package handler

import (
	"gateway/internal/filter"
	"gateway/internal/filter/gateway"
	"gateway/internal/filter/global"
	"gateway/internal/web"
	"gateway/logger"
)

func NewFilteringWebHandler() *FilteringWebHandler {

	return &FilteringWebHandler{
		globalFilters:  global.Filters,
		gatewayFilters: gateway.Filters,
	}
}

type FilteringWebHandler struct {
	globalFilters  []filter.OrderedFilter
	gatewayFilters []filter.OrderedFilter
}

func (h *FilteringWebHandler) Handle(exchange *web.ServerWebExchange) {
	logger.Logger.Debugf("start process filtering handler uri %s ", exchange.Request.URL.Path)

	var filters []filter.OrderedFilter
	for _, globalFilter := range h.globalFilters {
		filters = append(filters, globalFilter)
	}

	for _, gatewayFilter := range h.gatewayFilters {
		filters = append(filters, gatewayFilter)
	}

	filter.NewDefaultGatewayFilterChain(filters).Filter(exchange)
}
