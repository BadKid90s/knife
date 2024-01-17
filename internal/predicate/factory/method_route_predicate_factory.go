package factory

import (
	"gateway/internal/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/web"
	"gateway/logger"
)

type MethodRoutePredicateFactory struct {
	config *MethodPredicateConfig
}

func (f *MethodRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(), nil

}
func (f *MethodRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) error {
	args := getArgs(definition)
	f.config = &MethodPredicateConfig{
		methods: args,
	}
	return nil
}

func (f *MethodRoutePredicateFactory) apply() predicate.Predicate[*web.ServerWebExchange] {
	return &MethodPredicate[*web.ServerWebExchange]{
		methods: f.config.methods,
	}
}

// MethodPredicate
// 谓词信息
type MethodPredicate[T any] struct {
	predicate.DefaultPredicate[T]
	methods []string
}

func (p *MethodPredicate[T]) Apply(exchange *web.ServerWebExchange) bool {
	var result = false
	for _, method := range p.methods {
		if method == exchange.Request.Method {
			result = true
			break
		}
	}
	logger.Logger.Debugf("predicate apply success. result: %t, id: MethodPredicate ", result)
	return result
}

// MethodPredicateConfig
// 配置信息
type MethodPredicateConfig struct {
	methods []string
}

func (c *MethodPredicateConfig) ToString() string {
	return ""
}
