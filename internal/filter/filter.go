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

type OrderedFilter struct {
	Filter Filter
	Order  int16
}

func NewOrderedFilter(order int16, filter Filter) OrderedFilter {
	return OrderedFilter{
		Filter: filter,
		Order:  order,
	}
}
