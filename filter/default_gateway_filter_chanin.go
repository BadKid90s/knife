package filter

import (
	"gateway/web"
)

func NewDefaultGatewayFilterChain(filters []GatewayFilter) *DefaultGatewayFilterChain {
	return &DefaultGatewayFilterChain{
		filters: filters,
		index:   0,
	}
}

type DefaultGatewayFilterChain struct {
	filters []GatewayFilter

	index int
}

func (c *DefaultGatewayFilterChain) Filter(exchange *web.ServerWebExchange) {
	if c.index < len(c.filters) {
		filter := c.filters[c.index]
		c.index++
		filter.Filter(exchange, c)
	}
}
