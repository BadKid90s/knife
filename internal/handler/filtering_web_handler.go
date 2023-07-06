package handler

import (
	filter2 "gateway/internal/filter"
	"gateway/internal/filter/global"
	"gateway/internal/route"
	"gateway/internal/util"
	"gateway/internal/web"
	"gateway/logger"
	"sort"
)

func NewFilteringWebHandler() *FilteringWebHandler {
	return &FilteringWebHandler{}
}

type FilteringWebHandler struct {
	globalFilters []filter2.GatewayFilter
}

func (h *FilteringWebHandler) AddGlobalFilter(filter filter2.GatewayFilter) {
	h.globalFilters = append(h.globalFilters, filter)
}

func (h *FilteringWebHandler) Handle(exchange *web.ServerWebExchange) {
	logger.Logger.Debugf("start process filtering handler uri %s ", exchange.Request.URL.Path)

	r := exchange.Attributes[util.GatewayRouteAttr]
	gr, ok := r.(*route.Route)
	if !ok {
		return
	}

	gatewayFilters := gr.GatewayFilters

	globalFilter := global.Filters

	filters := append(gatewayFilters, globalFilter...)

	sort.Slice(filters, func(i, j int) bool {
		return filters[i].GetOrder() < filters[j].GetOrder()
	})

	filter2.NewDefaultGatewayFilterChain(filters).Filter(exchange)
}
