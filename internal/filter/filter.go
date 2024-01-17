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

type Constructor func(exchange *web.ServerWebExchange, chain Chain)

func (m Constructor) Filter(exchange *web.ServerWebExchange, chain Chain) {
	m(exchange, chain)
}

var GlobalFilters filtersDomain

func AddGlobalFilter(name string, order int16, filter Filter) {
	f := Info{
		order:  order,
		Name:   name,
		Filter: filter,
	}
	GlobalFilters.Filters = append(GlobalFilters.Filters, f)
}

var GatewayFilters filtersDomain

func AddGatewayFilters(name string, order int16, filter Filter) {
	f := Info{
		order:  order,
		Name:   name,
		Filter: filter,
	}
	GatewayFilters.Filters = append(GatewayFilters.Filters, f)
}

type filtersDomain struct {
	Filters []Info
}

type Info struct {
	Name   string
	Filter Filter
	order  int16
}
