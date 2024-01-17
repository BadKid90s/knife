package handler

import (
	"gateway/internal/filter"
	"gateway/internal/web"
	"gateway/logger"
)

func NewFilteringWebHandler(globalFilters []filter.GlobalFilter, gatewayFilters []filter.GatewayFilter) *FilteringWebHandler {
	filters := make([]filter.Filter, 0, len(globalFilters)+len(gatewayFilters))

	for _, f := range globalFilters {
		filters = append(filters, f)
	}

	for _, f := range gatewayFilters {
		filters = append(filters, f)
	}

	return &FilteringWebHandler{
		filters: filters,
	}
}

type FilteringWebHandler struct {
	filters []filter.Filter
}

func (h *FilteringWebHandler) Handle(exchange *web.ServerWebExchange) {
	logger.Logger.Debugf("start process filtering handler uri %s ", exchange.Request.URL.Path)

	filter.NewDefaultGatewayFilterChain(h.filters).Filter(exchange)
}
