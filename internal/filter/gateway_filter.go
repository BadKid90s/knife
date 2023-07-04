package filter

import (
	"gateway/internal/web"
)

type GatewayFilter interface {
	Filter(exchange *web.ServerWebExchange, chain GatewayFilterChain)
	GetOrder() int
}
