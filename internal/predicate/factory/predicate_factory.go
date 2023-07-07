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
	"After":   &AfterRoutePredicateFactory{},
	"Before":  &BeforeRoutePredicateFactory{},
	"Between": &BetweenRoutePredicateFactory{},
	"Cookie":  &CookieRoutePredicateFactory{},
	"Method":  &MethodRoutePredicateFactory{},
}

func getArgs(definition *definition.PredicateDefinition) []string {
	args := make([]string, 0)
	for _, value := range definition.Args {
		args = append(args, value)
	}
	return args
}
