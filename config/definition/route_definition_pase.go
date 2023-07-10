package definition

import (
	"gopkg.in/yaml.v2"
)

var RouteDefinitions = make([]*RouteDefinition, 0)

func ParseRouteConfig(buffer []byte) error {
	var config GatewayRoutesDefinition
	err := yaml.Unmarshal(buffer, &config)
	if err != nil {
		return err
	}

	routes := config.Routes
	for _, route := range routes {
		route.SetPredicateDefinitions()
	}
	RouteDefinitions = routes
	return nil
}
