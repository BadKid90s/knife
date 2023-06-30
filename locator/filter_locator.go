package locator

import (
	"gateway/definition"
	"gateway/filter"
	"log"
)

func getFilters(_ *definition.RouteDefinition) []filter.GatewayFilter {
	var fs = make([]filter.GatewayFilter, 0)
	log.Printf("completed loading routing filter, size [%d] \n", len(fs))
	return fs
}
