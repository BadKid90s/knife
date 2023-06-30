package locator

import (
	"fmt"
	"gateway/definition"
	"gateway/predicate"
	"gateway/predicate/factory"
	"gateway/web"
	"log"
)

// 组合谓词
func combinePredicates(routeDefinition *definition.RouteDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	predicates := routeDefinition.Predicates
	p, err := lookup(routeDefinition, predicates[0])
	if err != nil {
		return nil, err
	}
	for _, andPredicate := range predicates[1:] {
		found, err := lookup(routeDefinition, andPredicate)
		if err != nil {
			return nil, err
		}
		p = p.And(found)
	}
	log.Printf("completed loading routing predicates \n")
	return p, nil
}

func lookup(_ *definition.RouteDefinition, predicateDefinition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	f, ok := factory.PredicateFactories[predicateDefinition.Name]
	if !ok {
		return nil, fmt.Errorf("Unsupported predicate [%s] \n", predicateDefinition.Name)
	}
	return &predicate.DefaultPredicate[*web.ServerWebExchange]{
		Delegate: f.Apply(predicateDefinition),
	}, nil
}
