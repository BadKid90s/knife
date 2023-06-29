package route

import (
	"gateway/util"
	"gopkg.in/yaml.v2"
	"log"
)

var RouterDefinitions = make([]*Definition, 0)

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

	RouterDefinitions = config.Routes
	return nil
}
