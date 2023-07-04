package factory

import (
	"gateway/internal/predicate"
)

var PredicateFactories = map[string]predicate.RoutePredicateFactory{
	"Method": &MethodRoutePredicateFactory{},
	"After":  &AfterRoutePredicateFactory{},
}
