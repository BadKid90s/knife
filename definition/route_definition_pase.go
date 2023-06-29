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

	RouteDefinitions = config.Routes
	return nil
}
