package locator

import (
	"fmt"
	"gateway/config/definition"
	predicate2 "gateway/internal/predicate"
	"gateway/internal/predicate/factory"
	"gateway/internal/web"
	"log"
)

// 组合谓词
func combinePredicates(routeDefinition *definition.RouteDefinition) (predicate2.Predicate[*web.ServerWebExchange], error) {
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

func lookup(_ *definition.RouteDefinition, predicateDefinition *definition.PredicateDefinition) (predicate2.Predicate[*web.ServerWebExchange], error) {
	f, ok := factory.PredicateFactories[predicateDefinition.Name]
	if !ok {
		return nil, fmt.Errorf("Unsupported predicate [%s] \n", predicateDefinition.Name)
	}
	apply, err := f.Apply(predicateDefinition)
	if err != nil {
		return nil, err
	}
	return &predicate2.DefaultPredicate[*web.ServerWebExchange]{
		Delegate: apply,
	}, nil
}
