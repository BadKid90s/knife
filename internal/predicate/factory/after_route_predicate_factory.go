package factory

import (
	"gateway/internal/config"
	"gateway/internal/predicate"
	"gateway/internal/util"
	"gateway/internal/web"
	"gateway/logger"
	"time"
)

type AfterRoutePredicateFactory struct {
	config *AfterPredicateConfig
}

func (f *AfterRoutePredicateFactory) Apply(definition *config.PredicateConfiguration) (predicate.Predicate[*web.ServerWebExchange], error) {
	err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(), nil

}
func (f *AfterRoutePredicateFactory) parseConfig(definition *config.PredicateConfiguration) error {
	args := getArgs(definition)
	t, err := util.ParseTime(args[0])
	if err != nil {
		return err
	}
	f.config = &AfterPredicateConfig{
		time: *t,
	}
	return nil
}

func (f *AfterRoutePredicateFactory) apply() predicate.Predicate[*web.ServerWebExchange] {
	return &AfterPredicate[*web.ServerWebExchange]{
		time: f.config.time,
	}
}

type AfterPredicate[T any] struct {
	predicate.DefaultPredicate[T]
	time time.Time
}

func (p *AfterPredicate[T]) Apply(T) bool {
	result := time.Now().After(p.time)
	logger.Logger.Debugf("predicate apply success. result: %t, id: AfterPredicate", result)
	return result
}

type AfterPredicateConfig struct {
	time time.Time
}

func (c *AfterPredicateConfig) ToString() string {
	return ""
}
