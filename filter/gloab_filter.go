package filter

import "gateway/filter/global"

var GlobalFilter = []GatewayFilter{
	&global.RouteToRequestUrlFilter{},
	&global.WebClientHttpRoutingFilter{},
}
