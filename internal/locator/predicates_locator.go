package locator

import (
	"fmt"
	"gateway/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/predicate/factory"
	"gateway/internal/web"
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
	apply, err := f.Apply(predicateDefinition)
	if err != nil {
		return nil, err
	}
	if apply == nil {
		return nil, fmt.Errorf("an error occurred in building Predicate [%s]\n", predicateDefinition.Name)
	}
	return &predicate.DefaultPredicate[*web.ServerWebExchange]{
		Delegate: apply,
	}, nil
}
