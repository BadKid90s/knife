package handler

import (
	"fmt"
	"gateway/filters"
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
	sort.Slice(gatewayFilters, func(i, j int) bool {
		return gatewayFilters[i].GetOrder() < gatewayFilters[j].GetOrder()
	})

	filters.NewDefaultGatewayFilterChain(gatewayFilters).Filter(exchange)
}
