package gateway

import (
	"gateway/internal/config"
	"gateway/internal/filter"
	"gateway/internal/web"
	"gateway/logger"
	"net/url"
	"strings"
)

type StripPrefixFactory struct {
}

func (a *StripPrefixFactory) Apply(configuration *config.FilterConfiguration) filter.Filter {
	path := configuration.Args["path"]
	if path == nil {
		logger.Logger.TagLogger("stripPrefix").Fatalf("config path is null ")
	}
	prefix, ok := path.(string)
	if !ok {
		logger.Logger.TagLogger("stripPrefix").Fatalf("config path is not string type")
	}
	return filter.Constructor(func(exchange *web.ServerWebExchange, chain filter.Chain) {

		r := exchange.Request
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r.URL = new(url.URL)
			r.URL.Path = p
			r.URL.RawPath = rp
		}
		chain.Filter(exchange)
	})
}

func init() {
	AddFactory("StripPrefix", &StripPrefixFactory{})
}
