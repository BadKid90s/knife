package handler

import (
	"errors"
	"gateway/internal/locator"
	"gateway/internal/route"
	util2 "gateway/internal/util"
	web2 "gateway/internal/web"
	"gateway/logger"
)

func NewRoutePredicateHandlerMapping() *RoutePredicateHandlerMapping {
	return &RoutePredicateHandlerMapping{
		webHandler:    NewFilteringWebHandler(),
		routerLocator: locator.NewCachingRouteLocator(),
	}
}

type RoutePredicateHandlerMapping struct {
	webHandler    *FilteringWebHandler
	routerLocator locator.RouteLocator
}

func (r *RoutePredicateHandlerMapping) GetHandler(exchange *web2.ServerWebExchange) (Handler, error) {
	handler, err := r.getHandlerInternal(exchange)
	if err != nil {
		return nil, err
	}
	request := exchange.Request
	//判断是否支持跨域请求
	if handler != nil || util2.IsPreFlightRequest(request) {
		//corsProcessor.process(config, exchange)
	}
	return handler, nil
}

func (r RoutePredicateHandlerMapping) getHandlerInternal(exchange *web2.ServerWebExchange) (Handler, error) {
	//处理路由信息
	gatewayRoute, err := r.lookupRoute(exchange)
	if err != nil {
		return nil, err
	}
	exchange.Attributes[util2.GatewayRouteAttr] = gatewayRoute
	//返回
	return r.webHandler, nil
}

func (r *RoutePredicateHandlerMapping) lookupRoute(exchange *web2.ServerWebExchange) (*route.Route, error) {
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
