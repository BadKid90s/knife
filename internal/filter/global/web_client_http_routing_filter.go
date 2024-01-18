package global

import (
	"gateway/internal/filter"
	"gateway/internal/util"
	"gateway/internal/web"
)

type WebClientHttpRoutingFilter struct {
}

func (f *WebClientHttpRoutingFilter) Filter(exchange *web.ServerWebExchange, chain filter.Chain) {
	requestUrl := exchange.Attributes[util.GatewayRequestUrlAttr]
	reqUrl, ok := requestUrl.(string)
	if !ok {
		return
	}
	util.ServeReverseProxy(reqUrl, exchange.Write, exchange.Request)
	chain.Filter(exchange)
}

func init() {
	AddFilter("xxx", 1, &WebClientHttpRoutingFilter{})
}
