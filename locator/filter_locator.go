package locator

import (
	"gateway/definition"
	"gateway/filter"
	"log"
)

func getFilters(routeDefinition *definition.RouteDefinition) []filter.GatewayFilter {
	log.Printf("loading route filter route-id %s \n", routeDefinition.Id)
	var fs = make([]filter.GatewayFilter, 0)
	return fs
}
