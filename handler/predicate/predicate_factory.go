package predicate

import "gateway/web"

type RoutePredicateFactory interface {
	Apply(definition Definition) Predicate[*web.ServerWebExchange]
}
