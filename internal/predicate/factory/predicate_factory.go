package factory

import (
	"gateway/internal/config"
	"gateway/internal/predicate"
	"gateway/internal/web"
)

type RoutePredicateFactory interface {
	Apply(definition *config.PredicateConfiguration) (predicate.Predicate[*web.ServerWebExchange], error)
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

func getArgs(definition *config.PredicateConfiguration) []string {
	return definition.Args
}
