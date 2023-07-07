package factory

import (
	"errors"
	"gateway/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/web"
	"gateway/logger"
	"regexp"
)

type CookieRoutePredicateFactory struct {
	config *CookiePredicateConfig
}

func (f *CookieRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(), nil

}
func (f *CookieRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) error {
	args := getArgs(definition)
	if len(args) != 2 {
		return errors.New("need two datetime type params for CookieRoutePredicateFactory")
	}
	f.config = &CookiePredicateConfig{
		key:   args[0],
		value: args[1],
	}
	return nil
}

func (f *CookieRoutePredicateFactory) apply() predicate.Predicate[*web.ServerWebExchange] {
	return &CookiePredicate[*web.ServerWebExchange]{
		key:   f.config.key,
		value: f.config.value,
	}
}

// CookiePredicate
// 谓词信息
type CookiePredicate[T any] struct {
	predicate.DefaultPredicate[T]
	key   string
	value string
}

func (p *CookiePredicate[T]) Apply(exchange *web.ServerWebExchange) bool {
	var result = false
	cookie, err := exchange.Request.Cookie(p.key)
	if err != nil {
		logger.Logger.Debugf("cookiePredicate need name is %s param, but not found ", p.key)
		return false
	}
	result = regexp.MustCompile(p.value).MatchString(cookie.Value)
	logger.Logger.Debugf("predicate apply success. result: %t, id: CookiePredicate ", result)
	return result
}

// CookiePredicateConfig
// 配置信息
type CookiePredicateConfig struct {
	key   string
	value string
}

func (c *CookiePredicateConfig) ToString() string {
	return "CookiePredicateConfig"
}
