package filter

import (
	"gateway/internal/web"
)

type GatewayFilterChain interface {
	Filter(exchange *web.ServerWebExchange)
}
