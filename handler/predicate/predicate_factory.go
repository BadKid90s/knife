package predicate

import "gateway/web"

type RoutePredicateFactory interface {
	Apply() Predicate[*web.ServerWebExchange]
	NewConfig(definition Definition)
}

type AbstractRoutePredicateFactory struct {
}

func (f *AbstractRoutePredicateFactory) applyAsync() DefaultPredicate[*web.ServerWebExchange] {
	return DefaultPredicate[*web.ServerWebExchange]{
		Delegate: f.Apply(),
	}
}
func (f AbstractRoutePredicateFactory) Apply() Predicate[*web.ServerWebExchange] {
	return nil
}
