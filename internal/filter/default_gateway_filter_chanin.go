package filter

import (
	"gateway/internal/web"
	"gateway/logger"
	"reflect"
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
		logger.Logger.Debugf("global filter start. id: %s", reflect.TypeOf(filter).Elem().Name())
		filter.Filter(exchange, c)
		logger.Logger.Debugf("global filter success. id: %s", reflect.TypeOf(filter).Elem().Name())
	}
}
