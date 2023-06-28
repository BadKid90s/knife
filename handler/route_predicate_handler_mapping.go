package handler

import (
	"gateway/route"
	"gateway/util"
	"gateway/web"
)

func NewRoutePredicateHandlerMapping() *RoutePredicateHandlerMapping {
	return &RoutePredicateHandlerMapping{
		webHandler:    NewFilteringWebHandler(),
		routerLocator: route.NewDefinitionRouteLocator(),
	}
}

type RoutePredicateHandlerMapping struct {
	webHandler    *FilteringWebHandler
	routerLocator *route.DefinitionRouteLocator
}

func (r *RoutePredicateHandlerMapping) GetHandler(exchange *web.ServerWebExchange) web.Handler {
	handler := r.getHandlerInternal(exchange)
	request := exchange.Request
	//判断是否支持跨域请求
	if handler != nil || util.IsPreFlightRequest(request) {
		//corsProcessor.process(config, exchange)
	}
	return handler
}

func (r RoutePredicateHandlerMapping) getHandlerInternal(exchange *web.ServerWebExchange) web.Handler {
	//处理路由信息
	r.lookupRoute(exchange)
	//返回
	return r.webHandler
}

func (r *RoutePredicateHandlerMapping) lookupRoute(exchange *web.ServerWebExchange) *route.Route {
	for _, router := range r.routerLocator.Routes {
		//if router.Predicates.Apply(exchange) {
		return router
		//}
	}
	return nil
}

func init() {
	web.DispatcherHandlerConstant.AddHandler(NewRoutePredicateHandlerMapping())
}
