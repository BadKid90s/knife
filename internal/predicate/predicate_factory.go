package predicate

import (
	"gateway/config/definition"
	"gateway/internal/web"
)

type RoutePredicateFactory interface {
	Apply(definition *definition.PredicateDefinition) (Predicate[*web.ServerWebExchange], error)
}
