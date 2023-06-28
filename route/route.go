package route

import (
	"gateway/filters"
	"gateway/handler/predicate"
	"gateway/web"
)

type Route struct {
	Id             string
	Uri            string
	Order          string
	Predicates     predicate.Predicate[*web.ServerWebExchange]
	GatewayFilters []filters.GatewayFilter
}
