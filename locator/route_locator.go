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
