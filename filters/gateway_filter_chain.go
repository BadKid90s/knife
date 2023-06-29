package filters

import (
	"gateway/web"
)

func NewDefaultGatewayFilterChain(filters []GatewayFilter) *GatewayFilterChain {
	return &GatewayFilterChain{
		filters: filters,
		index:   0,
	}
}

type GatewayFilterChain struct {
	filters []GatewayFilter

	index int
}

func (c *GatewayFilterChain) Filter(exchange *web.ServerWebExchange) {
	if c.index < len(c.filters) {
		c.index++
		c.filters[c.index].Filter(exchange, c)
	}
}
