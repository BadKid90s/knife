package handler

import (
	"errors"
	"gateway/locator"
	"gateway/route"
	"gateway/util"
	"gateway/web"
	"log"
)

func NewRoutePredicateHandlerMapping() *RoutePredicateHandlerMapping {
	return &RoutePredicateHandlerMapping{
		webHandler:    NewFilteringWebHandler(),
		routerLocator: locator.NewDefinitionRouteLocator(),
	}
}

type RoutePredicateHandlerMapping struct {
	webHandler    *FilteringWebHandler
	routerLocator *locator.DefinitionRouteLocator
}

func (r *RoutePredicateHandlerMapping) GetHandler(exchange *web.ServerWebExchange) (web.Handler, error) {
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

func (r RoutePredicateHandlerMapping) getHandlerInternal(exchange *web.ServerWebExchange) (web.Handler, error) {
	//处理路由信息
	gatewayRoute, err := r.lookupRoute(exchange)
	if err != nil {
		return nil, err
	}
	exchange.Attributes[util.GatewayRouteAttr] = gatewayRoute
	//返回
	return r.webHandler, nil
}

func (r *RoutePredicateHandlerMapping) lookupRoute(exchange *web.ServerWebExchange) (*route.Route, error) {
	routes, err := r.routerLocator.GetRoutes()
	if err != nil {
		return nil, err
	}
	for _, r := range routes {
		if r.Predicates.Apply(exchange) {
			log.Printf("predicate route success [%s] \n", r.Id)
			return r, nil
		}
	}
	return nil, errors.New("no routing information matched \n")
}
