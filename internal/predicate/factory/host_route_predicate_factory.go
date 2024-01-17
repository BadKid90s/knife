package factory

import (
	"gateway/internal/config/definition"
	"gateway/internal/predicate"
	"gateway/internal/web"
	"gateway/logger"
	"github.com/vibrantbyte/go-antpath/antpath"
	"strings"
)

type HostRoutePredicateFactory struct {
	config *HostPredicateConfig
}

func (f *HostRoutePredicateFactory) Apply(definition *definition.PredicateDefinition) (predicate.Predicate[*web.ServerWebExchange], error) {
	err := f.parseConfig(definition)
	if err != nil {
		return nil, err
	}
	//return nil
	return f.apply(), nil

}
func (f *HostRoutePredicateFactory) parseConfig(definition *definition.PredicateDefinition) error {
	args := getArgs(definition)
	var hosts = make([]string, 0)
	for _, item := range args {
		split := strings.Split(item, ",")
		hosts = append(hosts, split...)
	}
	f.config = &HostPredicateConfig{
		hosts: hosts,
	}
	return nil
}

func (f *HostRoutePredicateFactory) apply() predicate.Predicate[*web.ServerWebExchange] {
	return &HostPredicate[*web.ServerWebExchange]{
		hosts: f.config.hosts,
	}
}

// HostPredicate
// 谓词信息
type HostPredicate[T any] struct {
	predicate.DefaultPredicate[T]
	hosts []string
}

func (p *HostPredicate[T]) Apply(exchange *web.ServerWebExchange) bool {
	var result = false
	requestHost := exchange.Request.Host
	matcher := antpath.New()
	for _, pattern := range p.hosts {
		match := matcher.Match(pattern, requestHost)
		if match {
			result = true
			break
		}
	}

	logger.Logger.Debugf("predicate apply success. result: %t, id: HostPredicate ", result)
	return result
}

// HostPredicateConfig
// 配置信息
type HostPredicateConfig struct {
	hosts []string
}

func (c *HostPredicateConfig) ToString() string {
	return "HostPredicateConfig"
}
