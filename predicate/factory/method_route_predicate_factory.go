package factory

import (
	"gateway/definition"
	"gateway/predicate"
	"gateway/web"
	"strings"
)

type MethodRoutePredicateFactory struct {
}

func (f *MethodRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	config, err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(config), nil

}
func (f *MethodRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) (*MethodPredicateConfig, error) {
	val := definition.Args["pattern"]
	methods := strings.Split(val, ",")
	return &MethodPredicateConfig{
		methods: methods,
	}, nil
}

func (f *MethodRoutePredicateFactory) apply(config *MethodPredicateConfig) predicate.Predicate[*web.ServerWebExchange] {
	return &MethodPredicate[*web.ServerWebExchange]{
		methods: config.methods,
	}
}

// MethodPredicate
// 谓词信息
type MethodPredicate[T any] struct {
	predicate.DefaultPredicate[T]
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
