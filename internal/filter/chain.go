package filter

import (
	"gateway/internal/web"
)

type Chain interface {
	Filter(exchange *web.ServerWebExchange)
}

func NewDefaultGatewayFilterChain(filters []OrderedFilter) *DefaultFilterChain {
	return &DefaultFilterChain{
		filters: filters,
		index:   0,
	}
}

type DefaultFilterChain struct {
	filters []OrderedFilter
	index   int
}

func (c *DefaultFilterChain) Filter(exchange *web.ServerWebExchange) {
	if c.index < len(c.filters) {
		orderedFilter := c.filters[c.index]
		c.index++
		orderedFilter.Filter.Filter(exchange, c)
	}
}
