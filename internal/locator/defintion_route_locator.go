package locator

import (
	"fmt"
	"gateway/config/definition"
	"gateway/internal/filter"
	"gateway/internal/predicate"
	"gateway/internal/predicate/factory"
	"gateway/internal/route"
	"gateway/internal/web"
	"log"
)

func NewDefinitionRouteLocator() *DefinitionRouteLocator {
	return &DefinitionRouteLocator{}
}

type DefinitionRouteLocator struct {
}

func (l *DefinitionRouteLocator) GetRoutes() ([]*route.Route, error) {
	routes := make([]*route.Route, 0)
	for _, routeDefinition := range definition.RouteDefinitions {
		r, err := l.ConvertToRoute(routeDefinition)
		if err != nil {
			return nil, err
		}
		routes = append(routes, r)
	}
	return routes, nil
}

func (l *DefinitionRouteLocator) ConvertToRoute(routeDefinition *definition.RouteDefinition) (*route.Route, error) {
	log.Printf("started covert route [%s] \n", routeDefinition.Id)
	predicates, err := combinePredicates(routeDefinition)
	if err != nil {
		return nil, err
	}
	gatewayFilters, err := getFilters(routeDefinition)
	if err != nil {
		return nil, err
	}
	return &route.Route{
		Id:             routeDefinition.Id,
		Uri:            routeDefinition.Uri,
		Order:          routeDefinition.Order,
		Predicates:     predicates,
		GatewayFilters: gatewayFilters,
	}, nil
}

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
func getFilters(_ *definition.RouteDefinition) ([]filter.GatewayFilter, error) {
	var fs = make([]filter.GatewayFilter, 0)
	log.Printf("completed loading routing filter, size [%d] \n", len(fs))
	return fs, nil
}
