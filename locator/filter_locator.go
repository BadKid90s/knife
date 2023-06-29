package locator

import (
	"gateway/definition"
	"gateway/filter"
	"gateway/filter/global"
)

func getFilters(routeDefinition *definition.RouteDefinition) []filter.GatewayFilter {
	fs := []filter.GatewayFilter{
		&global.RouteToRequestUrlFilter{},
		&global.WebClientHttpRoutingFilter{},
	}
	return fs
}
