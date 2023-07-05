package locator

import (
	"gateway/internal/route"
	"gateway/logger"
)

func NewCachingRouteLocator() *CachingRouteLocator {
	locator := &CachingRouteLocator{
		delegate: NewDefinitionRouteLocator(),
	}
	err := locator.fetch()
	if err != nil {
		logger.Logger.Fatalf(" failed to initialize route loader")
	}
	return locator
}

type CachingRouteLocator struct {
	delegate RouteLocator
	routes   []*route.Route
}

func (l *CachingRouteLocator) GetRoutes() ([]*route.Route, error) {
	return l.routes, nil
}
func (l *CachingRouteLocator) fetch() error {
	routes, err := l.delegate.GetRoutes()
	if err != nil {
		return err
	}
	l.routes = routes
	return nil
}
