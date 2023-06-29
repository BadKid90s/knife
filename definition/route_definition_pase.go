package definition

import (
	"gateway/util"
	"gopkg.in/yaml.v2"
	"log"
)

var RouteDefinitions = make([]*RouteDefinition, 0)

func ParseRouteConfig(configFile string) error {

	buffer, err := util.ReadConfigFile(configFile)
	if err != nil {
		return err
	}

	var config GatewayRoutesDefinition
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	RouteDefinitions = config.Routes
	return nil
}
