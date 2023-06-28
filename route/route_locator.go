package route

import (
	"gateway/config"
	"gateway/filters"
	"gateway/handler/predicate"
	"gateway/web"
)

var predicateFactories = map[string]predicate.RoutePredicateFactory{}

func NewDefinitionRouteLocator() *DefinitionRouteLocator {
	routes := make([]*Route, 0)
	for _, definition := range config.RouterDefinitions {
		route := ConvertToRoute(definition)
		routes = append(routes, route)
	}

	return &DefinitionRouteLocator{
		Routes: routes,
	}
}

type DefinitionRouteLocator struct {
	Routes []*Route
}

func ConvertToRoute(routeDefinition *Definition) *Route {
	predicates := combinePredicates(routeDefinition)
	gatewayFilters := getFilters(routeDefinition)
	return &Route{
		Id:             routeDefinition.Id,
		Uri:            routeDefinition.Uri,
		Order:          routeDefinition.Order,
		Predicates:     predicates,
		GatewayFilters: gatewayFilters,
	}
}

// 组合谓词
func combinePredicates(routeDefinition *Definition) predicate.Predicate[*web.ServerWebExchange] {
	predicates := routeDefinition.Predicates
	p := lookup(routeDefinition, predicates[0])
	for _, andPredicate := range predicates[1:] {
		found := lookup(routeDefinition, andPredicate)
		p = p.And(found)
	}
	return p
}

func lookup(routeDefinition *Definition, predicateDefinition predicate.Definition) predicate.Predicate[*web.ServerWebExchange] {
	factory := predicateFactories[predicateDefinition.Name]
	factory.NewConfig(predicateDefinition)
	return &predicate.DefaultPredicate[*web.ServerWebExchange]{
		Delegate: factory.Apply(),
	}
}

func getFilters(routeDefinition *Definition) []filters.GatewayFilter {
	return nil
}
