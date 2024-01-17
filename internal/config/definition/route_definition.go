package definition

import (
	"gateway/logger"
	"gopkg.in/yaml.v2"
	"strings"
)

type GatewayRoutesDefinition struct {
	Routes []*RouteDefinition `yaml:"routes"`
}

type RouteDefinition struct {
	Id                   string   `yaml:"id"`
	Uri                  string   `yaml:"uri"`
	Order                string   `yaml:"order"`
	Predicates           []string `yaml:"predicates"`
	PredicateDefinitions []*PredicateDefinition
}

type PredicateDefinition struct {
	Name string
	Args []string
}

func (r *RouteDefinition) SetPredicateDefinitions() {
	var press = make([]*PredicateDefinition, 0)
	for _, predStr := range r.Predicates {
		array := strings.Split(predStr, "=")
		if len(array) != 2 {
			logger.Logger.Fatalf("an error occurred while resolving the assertion %s with the route name being %s", predStr, r.Id)
		}
		predicate := &PredicateDefinition{
			Name: array[0],
			Args: strings.Split(array[1], ","),
		}
		press = append(press, predicate)
	}
	r.PredicateDefinitions = press
}

var RouteDefinitions = make([]*RouteDefinition, 0)

func ParseRouteConfig(buffer []byte) (*GatewayRoutesDefinition, error) {
	config := &GatewayRoutesDefinition{}
	err := yaml.Unmarshal(buffer, config)
	if err != nil {
		return nil, err
	}

	routes := config.Routes
	for _, route := range routes {
		route.SetPredicateDefinitions()
	}
	RouteDefinitions = routes
	return config, nil
}
