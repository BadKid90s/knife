package factory

import (
	"gateway/internal/config/definition"
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
	"Header":  &HeaderRoutePredicateFactory{},
	"Host":    &HostRoutePredicateFactory{},
	"Method":  &MethodRoutePredicateFactory{},
}

func getArgs(definition *definition.PredicateDefinition) []string {
	return definition.Args
}
