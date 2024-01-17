package handler

import (
	"errors"
	"gateway/internal/filter"
	"gateway/internal/locator"
	"gateway/internal/route"
	"gateway/internal/util"
	"gateway/internal/web"
	"gateway/logger"
)

func NewRoutePredicateHandlerMapping() *RoutePredicateHandlerMapping {
	return &RoutePredicateHandlerMapping{
		filterWebHandler: NewFilteringWebHandler(filter.GlobalFilters.Filters, filter.GatewayFilters.Filters),
		routerLocator:    locator.NewCachingRouteLocator(),
	}
}

type RoutePredicateHandlerMapping struct {
	filterWebHandler *FilteringWebHandler
	routerLocator    locator.RouteLocator
}

func (r *RoutePredicateHandlerMapping) GetHandler(exchange *web.ServerWebExchange) (Handler, error) {
	handler, err := r.getHandlerInternal(exchange)
	if err != nil {
		return nil, err
	}
	request := exchange.Request
	//判断是否支持跨域请求
	if handler != nil || util.IsPreFlightRequest(request) {
		//corsProcessor.process(config, exchange)
	}
	return handler, nil
}

func (r *RoutePredicateHandlerMapping) getHandlerInternal(exchange *web.ServerWebExchange) (Handler, error) {
	//处理路由信息
	gatewayRoute, err := r.lookupRoute(exchange)
	if err != nil {
		return nil, err
	}
	exchange.Attributes[util.GatewayRequestUrlAttr] = gatewayRoute.Uri
	//返回
	return r.filterWebHandler, nil
}

func (r *RoutePredicateHandlerMapping) lookupRoute(exchange *web.ServerWebExchange) (*route.Route, error) {
	routes, err := r.routerLocator.GetRoutes()
	if err != nil {
		return nil, err
	}
	for _, r := range routes {
		if r.Predicates.Apply(exchange) {
			logger.Logger.Debugf("predicate route success. route-id: %s ", r.Id)
			return r, nil
		}
	}
	return nil, errors.New("no routing information matched ")
}
