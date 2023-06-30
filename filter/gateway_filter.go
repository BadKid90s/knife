package filter

import "gateway/web"

type GatewayFilter interface {
	Filter(exchange *web.ServerWebExchange, chain GatewayFilterChain)
	GetOrder() int
}
