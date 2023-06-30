package factory

import "gateway/predicate"

var PredicateFactories = map[string]predicate.RoutePredicateFactory{
	"Method": &MethodRoutePredicateFactory{},
	"After":  &AfterRoutePredicateFactory{},
}
