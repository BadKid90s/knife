package route

import (
	"gateway/filter"
	"gateway/predicate"
	"gateway/web"
)

type Route struct {
	Id             string
	Uri            string
	Order          string
	Predicates     predicate.Predicate[*web.ServerWebExchange]
	GatewayFilters []filter.GatewayFilter
}
