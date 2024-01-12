package global

import (
	"gateway/internal/filter"
	"gateway/internal/util"
	"gateway/internal/web"
	"math"
)

type WebClientHttpRoutingFilter struct {
}

func (f *WebClientHttpRoutingFilter) Filter(exchange *web.ServerWebExchange, chain filter.GatewayFilterChain) {
	requestUrl := exchange.Attributes[util.GatewayRequestUrlAttr]
	reqUrl, ok := requestUrl.(string)
	if !ok {
		return
	}
	util.ServeReverseProxy(reqUrl, exchange.Write, exchange.Request)
	chain.Filter(exchange)
}

func (f *WebClientHttpRoutingFilter) GetOrder() int {
	return math.MinInt16
}
