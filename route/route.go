package route

import (
	"gateway/handler/predicate"
	"gateway/middleware/filters"
	"gateway/web"
)

type Route struct {
	Id             string
	Uri            string
	Order          string
	Predicates     predicate.Predicate[*web.ServerWebExchange]
	GatewayFilters []filters.GatewayFilter
}
