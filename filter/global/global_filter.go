package global

import "gateway/filter"

var Filters = []filter.GatewayFilter{
	&RouteToRequestUrlFilter{},
	&WebClientHttpRoutingFilter{},
}
