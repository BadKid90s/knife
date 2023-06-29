package handler

import (
	"fmt"
	"gateway/filter"
	"gateway/route"
	"gateway/util"
	"gateway/web"
	"sort"
)

func NewFilteringWebHandler() *FilteringWebHandler {
	return &FilteringWebHandler{}
}

type FilteringWebHandler struct {
}

func (h *FilteringWebHandler) Handle(exchange *web.ServerWebExchange) {
	fmt.Printf("FilteringWebHandler uri [%s] \n", exchange.Request.URL.Path)

	r := exchange.Attributes[util.GatewayRouteAttr]
	gr, ok := r.(*route.Route)
	if !ok {
		return
	}

	gatewayFilters := gr.GatewayFilters

	globalFilter := filter.GlobalFilter

	filters := append(gatewayFilters, globalFilter...)

	sort.Slice(filters, func(i, j int) bool {
		return filters[i].GetOrder() < filters[j].GetOrder()
	})

	filter.NewDefaultGatewayFilterChain(filters).Filter(exchange)
}
