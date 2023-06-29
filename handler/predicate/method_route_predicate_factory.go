package predicate

import (
	"gateway/web"
	"strings"
)

type MethodRoutePredicateFactory struct {
}

func (f *MethodRoutePredicateFactory) Apply(definition Definition) Predicate[*web.ServerWebExchange] {
	config := f.parseConfig(definition)
	//return nil
	return f.apply(config)

}
func (f *MethodRoutePredicateFactory) parseConfig(definition Definition) *MethodPredicateConfig {
	val := definition.Args["pattern"]
	methods := strings.Split(val, ",")
	return &MethodPredicateConfig{
		methods: methods,
	}
}

func (f *MethodRoutePredicateFactory) apply(config *MethodPredicateConfig) Predicate[*web.ServerWebExchange] {
	return &MethodPredicate[*web.ServerWebExchange]{
		methods: config.methods,
	}
}

// MethodPredicate
// 谓词信息
type MethodPredicate[T any] struct {
	DefaultPredicate[T]
	methods []string
}

func (p *MethodPredicate[T]) Apply(exchange *web.ServerWebExchange) bool {
	for _, method := range p.methods {
		if method == exchange.Request.Method {
			return true
		}
	}
	return false
}

// MethodPredicateConfig
// 配置信息
type MethodPredicateConfig struct {
	methods []string
}

func (c *MethodPredicateConfig) ToString() string {
	return ""
}
