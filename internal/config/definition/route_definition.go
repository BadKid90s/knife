package definition

import (
	"errors"
	"fmt"
	"gateway/logger"
	"gopkg.in/yaml.v2"
	"strings"
)

const Name = "name"
const Args = "args"

type GatewayRoutesDefinition struct {
	Routes []*RouteDefinition `yaml:"routes"`
}

type RouteDefinition struct {
	Id         string   `yaml:"id"`
	Uri        string   `yaml:"uri"`
	Order      string   `yaml:"order"`
	Predicates []string `yaml:"predicates"`
	Filters    []any    `yaml:"filters"`

	PredicateDefinitions []*PredicateDefinition
	FiltersDefinitions   []*FilterDefinition
}

type PredicateDefinition struct {
	Name string
	Args []string
}

type FilterDefinition struct {
	Name string
	Args map[any]any
}

func ParseRouteConfig(buffer []byte) (*GatewayRoutesDefinition, error) {
	config := &GatewayRoutesDefinition{}
	err := yaml.Unmarshal(buffer, config)
	if err != nil {
		return nil, err
	}

	routes := config.Routes
	for _, route := range routes {
		predicateDefinitions := parasPredicateDefinitions(route.Predicates)
		route.PredicateDefinitions = predicateDefinitions

		filtersDefinitions := parasFiltersDefinitions(route.Filters)
		route.FiltersDefinitions = filtersDefinitions
	}
	return config, nil
}

func parasFiltersDefinitions(filters []any) []*FilterDefinition {
	var fds = make([]*FilterDefinition, 0)

	for _, item := range filters {

		fd := &FilterDefinition{
			Args: make(map[any]any),
		}
		switch v := item.(type) {
		case map[any]any:
			value, ok := v[Name].(string)
			if ok {
				fd.Name = value
			} else {
				logger.Logger.Fatalf("an error occurred while resolving the assertion %s ", v)
			}
			m, ok := v[Args].(map[any]any)
			if ok {
				fd.Args = m
			} else {
				logger.Logger.Fatalf("an error occurred while resolving the assertion %s ", v)
			}

		case string:
			array := strings.Split(v, "=")
			if len(array) != 2 {
				logger.Logger.Fatalf("an error occurred while resolving the assertion %s ", v)
			}
			fd.Name = array[0]
			for i, it := range strings.Split(array[1], ",") {
				fd.Args[generatorKey(i)] = it
			}
		default:
			fmt.Printf("x has unknown type %s", v)
		}

		fds = append(fds, fd)
	}

	return fds
}

func parasPredicateDefinitions(predicates []string) []*PredicateDefinition {
	var press = make([]*PredicateDefinition, 0)
	for _, predStr := range predicates {
		array := strings.Split(predStr, "=")
		if len(array) != 2 {
			logger.Logger.Fatalf("an error occurred while resolving the assertion %s ", predStr)
		}
		predicate := &PredicateDefinition{
			Name: array[0],
			Args: strings.Split(array[1], ","),
		}
		press = append(press, predicate)
	}
	return press
}

func generatorKey(i int) string {
	return fmt.Sprintf("_key_%d", i)
}

func convString(val any) (string, error) {
	str, ok := val.(string)
	if ok {
		return str, nil
	}
	return "", errors.New("conv string error")

}
