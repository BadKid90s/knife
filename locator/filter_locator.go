package locator

import (
	"gateway/definition"
	"gateway/filter"
	"log"
)

func getFilters(routeDefinition *definition.RouteDefinition) []filter.GatewayFilter {
	log.Printf("RouteDefinition-%s 加载配置的过滤器。\n", routeDefinition.Id)
	var fs = make([]filter.GatewayFilter, 0)
	return fs
}
