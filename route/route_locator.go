package route

import (
	"gateway/filters"
	"gateway/handler/predicate"
	"gateway/web"
)

var predicateFactories = map[string]predicate.RoutePredicateFactory{
	"Method": &predicate.MethodRoutePredicateFactory{},
	"After":  &predicate.AfterRoutePredicateFactory{},
}

func NewDefinitionRouteLocator() *DefinitionRouteLocator {
	return &DefinitionRouteLocator{}
}

type DefinitionRouteLocator struct {
}

func (l *DefinitionRouteLocator) GetRoutes() []*Route {
	routes := make([]*Route, 0)
	for _, definition := range RouterDefinitions {
		route := ConvertToRoute(definition)
		routes = append(routes, route)
	}
	return routes
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
	return &predicate.DefaultPredicate[*web.ServerWebExchange]{
		Delegate: factory.Apply(predicateDefinition),
	}
}

func getFilters(routeDefinition *Definition) []filters.GatewayFilter {
	fs := []filters.GatewayFilter{
		&filters.WebClientHttpRoutingFilter{},
	}
	return fs
}
