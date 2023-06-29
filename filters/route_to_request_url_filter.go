package filters

import (
	"gateway/route"
	"gateway/util"
	"gateway/web"
	"math"
)

type RouteToRequestUrlFilter struct {
}

func (f *RouteToRequestUrlFilter) Filter(exchange *web.ServerWebExchange, chain *GatewayFilterChain) {
	routeAttr := exchange.Attributes[util.GatewayRouteAttr]
	gatewayRoute, ok := routeAttr.(*route.Route)
	if !ok {
		return
	}
	uri := gatewayRoute.Uri
	exchange.Attributes[util.GatewayRequestUrlAttr] = uri
	chain.Filter(exchange)
}

func (f *RouteToRequestUrlFilter) GetOrder() int {
	return math.MinInt16
}
