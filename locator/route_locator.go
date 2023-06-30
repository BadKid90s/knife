package locator

import (
	"gateway/definition"
	"gateway/route"
	"log"
)

func NewDefinitionRouteLocator() *DefinitionRouteLocator {
	return &DefinitionRouteLocator{}
}

type DefinitionRouteLocator struct {
}

func (l *DefinitionRouteLocator) GetRoutes() []*route.Route {
	routes := make([]*route.Route, 0)
	for _, routeDefinition := range definition.RouteDefinitions {
		r := l.ConvertToRoute(routeDefinition)
		routes = append(routes, r)
	}
	return routes
}

func (l *DefinitionRouteLocator) ConvertToRoute(routeDefinition *definition.RouteDefinition) *route.Route {
	log.Printf("started covert route  [%s] \n", routeDefinition.Id)
	predicates := combinePredicates(routeDefinition)
	gatewayFilters := getFilters(routeDefinition)
	return &route.Route{
		Id:             routeDefinition.Id,
		Uri:            routeDefinition.Uri,
		Order:          routeDefinition.Order,
		Predicates:     predicates,
		GatewayFilters: gatewayFilters,
	}
}
