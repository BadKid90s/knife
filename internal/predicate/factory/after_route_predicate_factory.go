package factory

import (
	"gateway/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/util"
	"gateway/internal/web"
	"time"
)

type AfterRoutePredicateFactory struct {
	config *AfterPredicateConfig
}

func (f *AfterRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(), nil

}
func (f *AfterRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) error {
	args := make([]string, 0)
	for _, value := range definition.Args {
		args = append(args, value)
	}
	t := util.ParseTime(args[0])
	f.config = &AfterPredicateConfig{
		time: t,
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
	return time.Now().After(p.time)
}

type AfterPredicateConfig struct {
	time time.Time
}

func (c *AfterPredicateConfig) ToString() string {
	return ""
}
