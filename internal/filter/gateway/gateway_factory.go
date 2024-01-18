package gateway

import (
	"gateway/internal/config"
	"gateway/internal/filter"
)

type Factory interface {
	Apply(configuration *config.FilterConfiguration) filter.Filter
	GetOrder() int
}

var Factories = make(map[string]Factory)

func AddFactory(name string, f Factory) {
	Factories[name] = f
}
