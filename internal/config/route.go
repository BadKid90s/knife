package config

import (
	"gateway/logger"
	"strings"
)

type GatewayRoutesConfiguration struct {
	Routes []*RouteConfiguration
}

type PredicateConfiguration struct {
	Name string
	Args []string
}

type RouteConfiguration struct {
	Id                   string
	Uri                  string
	Order                string
	Predicates           []string
	PredicateDefinitions []*PredicateConfiguration
}

func (r *RouteConfiguration) SetPredicateDefinitions() {
	var press = make([]*PredicateConfiguration, 0)
	for _, predStr := range r.Predicates {
		array := strings.Split(predStr, "=")
		if len(array) != 2 {
			logger.Logger.Fatalf("an error occurred while resolving the assertion %s with the route name being %s", predStr, r.Id)
		}
		predicate := &PredicateConfiguration{
			Name: array[0],
			Args: strings.Split(array[1], ","),
		}
		press = append(press, predicate)
	}
	r.PredicateDefinitions = press
}
