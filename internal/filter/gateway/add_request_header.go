package gateway

import (
	"gateway/internal/config"
	"gateway/internal/filter"
	"gateway/internal/web"
)

type AddRequestHeaderFactory struct {
}

func (a *AddRequestHeaderFactory) GetOrder() int {
	return 1
}

func (a *AddRequestHeaderFactory) Apply(configuration *config.FilterConfiguration) filter.Filter {
	return filter.Constructor(func(exchange *web.ServerWebExchange, chain filter.Chain) {

	})
}

func init() {
	AddFactory("AddRequestHeader", &AddRequestHeaderFactory{})
}
