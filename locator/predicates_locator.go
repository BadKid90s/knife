package locator

import (
	"gateway/definition"
	"gateway/predicate"
	"gateway/predicate/factory"
	"gateway/web"
)

// 组合谓词
func combinePredicates(routeDefinition *definition.RouteDefinition) predicate.Predicate[*web.ServerWebExchange] {
	predicates := routeDefinition.Predicates
	p := lookup(routeDefinition, predicates[0])
	for _, andPredicate := range predicates[1:] {
		found := lookup(routeDefinition, andPredicate)
		p = p.And(found)
	}
	return p
}

func lookup(_ *definition.RouteDefinition, predicateDefinition *definition.PredicateDefinition) predicate.Predicate[*web.ServerWebExchange] {
	f := factory.PredicateFactories[predicateDefinition.Name]
	return &predicate.DefaultPredicate[*web.ServerWebExchange]{
		Delegate: f.Apply(predicateDefinition),
	}
}
