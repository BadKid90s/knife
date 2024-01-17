package factory

import (
	"gateway/internal/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/util"
	"gateway/internal/web"
	"gateway/logger"
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
	t, err := util.ParseTime(args[0])
	if err != nil {
		return err
	}
	f.config = &BeforePredicateConfig{
		time: *t,
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
	result := time.Now().Before(p.time)
	logger.Logger.Debugf("predicate apply success. result: %t, id: BeforePredicate ", result)
	return result
}

type BeforePredicateConfig struct {
	time time.Time
}

func (c *BeforePredicateConfig) ToString() string {
	return "BeforePredicateConfig"
}
