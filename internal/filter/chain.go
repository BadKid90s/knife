package filter

import (
	"gateway/internal/web"
)

type Chain interface {
	Filter(exchange *web.ServerWebExchange)
}

func NewDefaultGatewayFilterChain(filters []Filter) *DefaultFilterChain {
	return &DefaultFilterChain{
		filters: filters,
		index:   0,
	}
}

type DefaultFilterChain struct {
	filters []Filter
	index   int
}

func (c *DefaultFilterChain) Filter(exchange *web.ServerWebExchange) {
	if c.index < len(c.filters) {
		filter := c.filters[c.index]
		c.index++
		//logger.Logger.Debugf("global filter start. id: %s", reflect.TypeOf(filter).Elem().Name())
		filter.Filter(exchange, c)
		//logger.Logger.Debugf("global filter success. id: %s", reflect.TypeOf(filter).Elem().Name())
	}
}
