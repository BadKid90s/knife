package global

import (
	"gateway/internal/filter"
)

var Filters = []filter.GatewayFilter{
	&RouteToRequestUrlFilter{},
	&WebClientHttpRoutingFilter{},
}
