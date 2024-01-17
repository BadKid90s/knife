package handler

import (
	"gateway/internal/filter"
	"gateway/internal/web"
	"gateway/logger"
)

func NewFilteringWebHandler(globalFilters []filter.Info, gatewayFilters []filter.Info) *FilteringWebHandler {

	return &FilteringWebHandler{
		globalFilters:  globalFilters,
		gatewayFilters: gatewayFilters,
	}
}

type FilteringWebHandler struct {
	globalFilters  []filter.Info
	gatewayFilters []filter.Info
}

func (h *FilteringWebHandler) Handle(exchange *web.ServerWebExchange) {
	logger.Logger.Debugf("start process filtering handler uri %s ", exchange.Request.URL.Path)

	var filters []filter.Filter
	for _, globalFilter := range h.globalFilters {
		filters = append(filters, globalFilter.Filter)
	}

	for _, gatewayFilter := range h.gatewayFilters {
		filters = append(filters, gatewayFilter.Filter)
	}

	filter.NewDefaultGatewayFilterChain(filters).Filter(exchange)
}
