package filter

import (
	"gateway/internal/web"
)

type Filter interface {
	Filter(exchange *web.ServerWebExchange, chain Chain)
}

type Constructor func(exchange *web.ServerWebExchange, chain Chain)

func (m Constructor) Filter(exchange *web.ServerWebExchange, chain Chain) {
	m(exchange, chain)
}

//
//var GatewayFilters []OrderedFilter
//
//func AddGatewayFilters(name string, order int16, filter Filter) {
//	f := OrderedFilter{
//		order:  order,
//		Name:   name,
//		Filter: filter,
//	}
//	GatewayFilters = append(GatewayFilters, f)
//}

type OrderedFilter struct {
	Name   string
	Filter Filter
	Order  int16
}

func NewOrderedFilter(name string, order int16, filter Filter) OrderedFilter {
	return OrderedFilter{
		Name:   name,
		Filter: filter,
		Order:  order,
	}
}
