package handler

import (
	"fmt"
	"gateway/filter"
	"gateway/filter/global"
	"gateway/route"
	"gateway/util"
	"gateway/web"
	"sort"
)

func NewFilteringWebHandler() *FilteringWebHandler {
	return &FilteringWebHandler{}
}

type FilteringWebHandler struct {
	globalFilters []filter.GatewayFilter
}

func (h *FilteringWebHandler) AddGlobalFilter(filter filter.GatewayFilter) {
	h.globalFilters = append(h.globalFilters, filter)
}

func (h *FilteringWebHandler) Handle(exchange *web.ServerWebExchange) {
	fmt.Printf("FilteringWebHandler uri [%s] \n", exchange.Request.URL.Path)

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

	filter.NewDefaultGatewayFilterChain(filters).Filter(exchange)
}
