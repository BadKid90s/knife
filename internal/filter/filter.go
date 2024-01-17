package filter

import (
	"gateway/internal/web"
)

type Filter interface {
	Filter(exchange *web.ServerWebExchange, chain Chain)
}
type GlobalFilter interface {
	Filter
}
type GatewayFilter interface {
	Filter
}

var GatewayFilters []GatewayFilter
var GlobalFilters []GlobalFilter

type Constructor func(exchange *web.ServerWebExchange, chain Chain)

func (m Constructor) Filter(exchange *web.ServerWebExchange, chain Chain) {
	m(exchange, chain)
}
func AddGlobalFilter(order int16, filter Filter) {
	GlobalFilters = append(GlobalFilters, filter)
}
