package factory

import (
	"errors"
	"gateway/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/util"
	"gateway/internal/web"
	"gateway/logger"
	"time"
)

type BetweenRoutePredicateFactory struct {
	config *BetweenPredicateConfig
}

func (f *BetweenRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(), nil

}
func (f *BetweenRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) error {
	args := getArgs(definition)
	if len(args) != 2 {
		return errors.New("need two datetime type params for betweenRoutePredicate")
	}
	startTime, err := util.ParseTime(args[0])
	if err != nil {
		return err
	}
	endTime, err := util.ParseTime(args[1])
	if err != nil {
		return err
	}
	f.config = &BetweenPredicateConfig{
		startTime: *startTime,
		endTime:   *endTime,
	}
	return nil
}

func (f *BetweenRoutePredicateFactory) apply() predicate.Predicate[*web.ServerWebExchange] {
	return &BetweenPredicate[*web.ServerWebExchange]{
		startTime: f.config.startTime,
		endTime:   f.config.endTime,
	}

}

type BetweenPredicate[T any] struct {
	predicate.DefaultPredicate[T]
	startTime time.Time
	endTime   time.Time
}

func (p *BetweenPredicate[T]) Apply(T) bool {
	now := time.Now()
	result := now.After(p.startTime) && now.Before(p.endTime)
	logger.Logger.Debugf("predicate apply success. result: %t,  id: BetweenPredicate ", result)
	return result
}

type BetweenPredicateConfig struct {
	startTime time.Time
	endTime   time.Time
}

func (c *BetweenPredicateConfig) ToString() string {
	return "BetweenPredicateConfig"
}
