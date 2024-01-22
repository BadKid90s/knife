package gateway

import (
	"gateway/internal/config"
	"gateway/internal/filter"
	"gateway/internal/web"
	"gateway/logger"
)

type AddRequestHeaderFactory struct {
}

func (a *AddRequestHeaderFactory) Apply(configuration *config.FilterConfiguration) filter.Filter {

	args := configuration.Args
	var header string
	var headerValue string
	for i, value := range args {
		header = i
		val, ok := value.(string)
		if !ok {
			logger.Logger.TagLogger("addRequestHeaderFactory").Fatalf("config value is not string type, key:%s", i)
		}
		headerValue = val
	}
	return filter.Constructor(func(exchange *web.ServerWebExchange, chain filter.Chain) {
		exchange.Request.Header.Add(header, headerValue)
		chain.Filter(exchange)
	})
}

func init() {
	AddFactory("AddRequestHeader", &AddRequestHeaderFactory{})
}
