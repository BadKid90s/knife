package locator

import (
	"gateway/config/definition"
	"gateway/internal/filter"
	"log"
)

func getFilters(_ *definition.RouteDefinition) ([]filter.GatewayFilter, error) {
	var fs = make([]filter.GatewayFilter, 0)
	log.Printf("completed loading routing filter, size [%d] \n", len(fs))
	return fs, nil
}
