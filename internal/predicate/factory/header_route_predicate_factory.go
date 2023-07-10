package factory

import (
	"errors"
	"gateway/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/web"
	"gateway/logger"
	"regexp"
)

type HeaderRoutePredicateFactory struct {
	config *HeaderPredicateConfig
}

func (f *HeaderRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(), nil

}
func (f *HeaderRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) error {
	args := getArgs(definition)
	if len(args) != 2 {
		return errors.New("need two datetime type params for HeaderRoutePredicateFactory")
	}
	f.config = &HeaderPredicateConfig{
		key:   args[0],
		value: args[1],
	}
	return nil
}

func (f *HeaderRoutePredicateFactory) apply() predicate.Predicate[*web.ServerWebExchange] {
	return &HeaderPredicate[*web.ServerWebExchange]{
		key:   f.config.key,
		value: f.config.value,
	}
}

// HeaderPredicate
// 谓词信息
type HeaderPredicate[T any] struct {
	predicate.DefaultPredicate[T]
	key   string
	value string
}

func (p *HeaderPredicate[T]) Apply(exchange *web.ServerWebExchange) bool {
	var result = false
	value := exchange.Request.Header.Get(p.key)
	if len(value) == 0 {
		logger.Logger.Debugf("headerPredicate need name is %s param, but not found ", p.key)
	} else {
		result = regexp.MustCompile(p.value).MatchString(value)
	}
	logger.Logger.Debugf("predicate apply success. result: %t, id: HeaderPredicate ", result)
	return result
}

// HeaderPredicateConfig
// 配置信息
type HeaderPredicateConfig struct {
	key   string
	value string
}

func (c *HeaderPredicateConfig) ToString() string {
	return "HeaderPredicateConfig"
}
