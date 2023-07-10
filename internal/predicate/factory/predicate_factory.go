package factory

import (
	"gateway/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/web"
)

type RoutePredicateFactory interface {
	Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error)
}

var PredicateFactories = map[string]RoutePredicateFactory{
	"Method": &MethodRoutePredicateFactory{},
	"After":  &AfterRoutePredicateFactory{},
	"Before": &BeforeRoutePredicateFactory{},
}

func getArgs(definition *definition.PredicateDefinition) []string {
	return definition.Args
}
