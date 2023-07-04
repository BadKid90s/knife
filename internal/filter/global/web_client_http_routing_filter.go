package global

import (
	"gateway/internal/filter"
	util2 "gateway/internal/util"
	"gateway/internal/web"
	"math"
	"net/http/httputil"
	"net/url"
)

type WebClientHttpRoutingFilter struct {
}

func (f *WebClientHttpRoutingFilter) Filter(exchange *web.ServerWebExchange, chain filter.GatewayFilterChain) {
	requestUrl := exchange.Attributes[util2.GatewayRequestUrlAttr]
	reqUrl, ok := requestUrl.(string)
	if !ok {
		return
	}
	targetUrl, err := url.Parse(reqUrl)
	if err != nil {
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.ErrorLog = util2.NewHttpLogger()
	proxy.ServeHTTP(exchange.Write, exchange.Request)
	chain.Filter(exchange)
}

func (f *WebClientHttpRoutingFilter) GetOrder() int {
	return math.MinInt16
}
