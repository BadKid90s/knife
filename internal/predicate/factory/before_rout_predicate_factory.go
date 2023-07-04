package factory

import (
	"gateway/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/util"
	"gateway/internal/web"
	"time"
)

type BeforeRoutePredicateFactory struct {
	config *BeforePredicateConfig
}

func (f *BeforeRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(), nil

}
func (f *BeforeRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) error {
	args := getArgs(definition)
	t := util.ParseTime(args[0])
	f.config = &BeforePredicateConfig{
		time: t,
	}
	return nil
}

func (f *BeforeRoutePredicateFactory) apply() predicate.Predicate[*web.ServerWebExchange] {
	return &BeforePredicate[*web.ServerWebExchange]{
		time: f.config.time,
	}

}

type BeforePredicate[T any] struct {
	predicate.DefaultPredicate[T]
	time time.Time
}

func (p *BeforePredicate[T]) Apply(T) bool {
	return time.Now().Before(p.time)
}

type BeforePredicateConfig struct {
	time time.Time
}

func (c *BeforePredicateConfig) ToString() string {
	return "BeforePredicateConfig"
}
