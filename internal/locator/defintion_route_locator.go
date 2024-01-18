package locator

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/filter"
	"gateway/internal/filter/gateway"
	"gateway/internal/predicate"
	"gateway/internal/predicate/factory"
	"gateway/internal/route"
	"gateway/internal/web"
	"gateway/logger"
)

func NewDefinitionRouteLocator(routesConfiguration *config.GatewayRoutesConfiguration) *DefinitionRouteLocator {
	return &DefinitionRouteLocator{
		routesConfiguration: routesConfiguration,
	}
}

type DefinitionRouteLocator struct {
	routesConfiguration *config.GatewayRoutesConfiguration
}

func (l *DefinitionRouteLocator) GetRoutes() ([]*route.Route, error) {
	routes := make([]*route.Route, 0)
	for _, r := range l.routesConfiguration.Routes {
		r, err := l.ConvertToRoute(r)
		if err != nil {
			return nil, err
		}
		routes = append(routes, r)
	}
	return routes, nil
}

func (l *DefinitionRouteLocator) ConvertToRoute(routeDefinition *config.RouteConfiguration) (*route.Route, error) {
	logger.Logger.TagLogger("locator").Debugf("started covert route, route-id: %s", routeDefinition.Id)
	predicates, err := combinePredicates(routeDefinition)
	if err != nil {
		return nil, err
	}
	filters, err := getFilters(routeDefinition)
	if err != nil {
		return nil, err
	}
	return &route.Route{
		Id:         routeDefinition.Id,
		Uri:        routeDefinition.Uri,
		Order:      routeDefinition.Order,
		Predicates: predicates,
		Filters:    filters,
	}, nil
}

// 组合谓词
func combinePredicates(routeDefinition *config.RouteConfiguration) (predicate.Predicate[*web.ServerWebExchange], error) {
	predicates := routeDefinition.PredicateConfiguration
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
		logger.Logger.TagLogger("locator").Debugf("completed loading routing predicates，total: %d", len(predicates))
		return p, nil
	}
	return &predicate.NullableDefaultPredicate[*web.ServerWebExchange]{}, nil
}

func lookup(_ *config.RouteConfiguration, predicateDefinition *config.PredicateConfiguration) (predicate.Predicate[*web.ServerWebExchange], error) {
	f, ok := factory.PredicateFactories[predicateDefinition.Name]
	if !ok {
		return nil, fmt.Errorf("locator: Unsupported predicate [%s] ", predicateDefinition.Name)
	}
	apply, err := f.Apply(predicateDefinition)
	if err != nil {
		return nil, err
	}
	if apply == nil {
		return nil, fmt.Errorf("locator: an error occurred in building Predicate [%s] ", predicateDefinition.Name)
	}
	return apply, nil
}
func getFilters(router *config.RouteConfiguration) ([]filter.OrderedFilter, error) {
	var gatewayFilterList []filter.OrderedFilter

	for _, configuration := range router.FilterConfiguration {
		gatewayFactory := gateway.Factories[configuration.Name]
		if gatewayFactory != nil {
			order := gatewayFactory.GetOrder()
			gatewayFilter := gatewayFactory.Apply(configuration)
			gatewayFilterList = append(gatewayFilterList, filter.NewOrderedFilter(order, gatewayFilter))
		} else {
			logger.Logger.TagLogger("locator").Warnf("filter configuration is error,not matched gateway filter, filter name: %s", configuration.Name)
		}
	}
	logger.Logger.TagLogger("locator").Debugf("completed loading routing filter, total: %d ", len(gatewayFilterList))
	return gatewayFilterList, nil
}
