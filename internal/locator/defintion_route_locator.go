package locator

import (
	"fmt"
	definition2 "gateway/internal/config/definition"
	"gateway/internal/filter"
	"gateway/internal/predicate"
	"gateway/internal/predicate/factory"
	"gateway/internal/route"
	"gateway/internal/web"
	"gateway/logger"
)

func NewDefinitionRouteLocator() *DefinitionRouteLocator {
	return &DefinitionRouteLocator{}
}

type DefinitionRouteLocator struct {
}

func (l *DefinitionRouteLocator) GetRoutes() ([]*route.Route, error) {
	routes := make([]*route.Route, 0)
	for _, routeDefinition := range definition2.RouteDefinitions {
		r, err := l.ConvertToRoute(routeDefinition)
		if err != nil {
			return nil, err
		}
		routes = append(routes, r)
	}
	return routes, nil
}

func (l *DefinitionRouteLocator) ConvertToRoute(routeDefinition *definition2.RouteDefinition) (*route.Route, error) {
	logger.Logger.Debugf("started covert route, route-id: %s", routeDefinition.Id)
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
func combinePredicates(routeDefinition *definition2.RouteDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	predicates := routeDefinition.PredicateDefinitions
	if len(predicates) > 0 {
		p, err := lookup(routeDefinition, predicates[0])
		if err != nil {
			return nil, err
		}
		for _, andPredicate := range predicates[1:] {
			found, err := lookup(routeDefinition, andPredicate)
			if err != nil {
				return nil, err
			}
			p = &predicate.AndPredicate[*web.ServerWebExchange]{
				Left:  p,
				Right: found,
			}
		}
		logger.Logger.Debugf("completed loading routing predicates，total: %d", len(predicates))
		return p, nil
	}
	return &predicate.NullableDefaultPredicate[*web.ServerWebExchange]{}, nil
}

func lookup(_ *definition2.RouteDefinition, predicateDefinition *definition2.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	f, ok := factory.PredicateFactories[predicateDefinition.Name]
	if !ok {
		return nil, fmt.Errorf("Unsupported predicate [%s] ", predicateDefinition.Name)
	}
	apply, err := f.Apply(predicateDefinition)
	if err != nil {
		return nil, err
	}
	if apply == nil {
		return nil, fmt.Errorf("an error occurred in building Predicate [%s] ", predicateDefinition.Name)
	}
	return apply, nil
}

func getFilters(_ *definition2.RouteDefinition) ([]filter.GatewayFilter, error) {
	var fs = make([]filter.GatewayFilter, 0)
	logger.Logger.Debugf("completed loading routing filter, total: %d ", len(fs))
	return fs, nil
}
