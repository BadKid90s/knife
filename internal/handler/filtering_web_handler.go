package handler

import (
	"gateway/internal/filter"
	"gateway/internal/filter/global"
	"gateway/internal/route"
	"gateway/internal/util"
	"gateway/internal/web"
	"gateway/logger"
)

func NewFilteringWebHandler() *FilteringWebHandler {

	return &FilteringWebHandler{
		globalFilters: global.Filters,
	}
}

type FilteringWebHandler struct {
	globalFilters []filter.OrderedFilter
}

func (h *FilteringWebHandler) Handle(exchange *web.ServerWebExchange) {
	logger.Logger.Debugf("start process filtering handler uri %s ", exchange.Request.URL.Path)

	var filters []filter.OrderedFilter
	for _, globalFilter := range h.globalFilters {
		filters = append(filters, globalFilter)
	}

	r := exchange.Attributes[util.GatewayRouteAttr]
	router, ok := r.(route.Route)
	if !ok {
		return
	}
	for _, gatewayFilter := range router.Filters {
		filters = append(filters, gatewayFilter)
	}
	filter.NewDefaultGatewayFilterChain(filters).Filter(exchange)
}
