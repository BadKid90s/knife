package predicate

import (
	"gateway/definition"
	"gateway/web"
)

type RoutePredicateFactory interface {
	Apply(definition *definition.PredicateDefinition) (Predicate[*web.ServerWebExchange], error)
}
