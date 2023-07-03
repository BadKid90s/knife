package route

import (
	"gateway/internal/filter"
	"gateway/internal/predicate"
	"gateway/internal/web"
)

type Route struct {
	Id             string
	Uri            string
	Order          string
	Predicates     predicate.Predicate[*web.ServerWebExchange]
	GatewayFilters []filter.GatewayFilter
}
