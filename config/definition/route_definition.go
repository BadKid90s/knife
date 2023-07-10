package definition

import (
	"gateway/logger"
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
