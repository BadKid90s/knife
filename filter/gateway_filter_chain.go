package filter

import "gateway/web"

type GatewayFilterChain interface {
	Filter(exchange *web.ServerWebExchange)
}
