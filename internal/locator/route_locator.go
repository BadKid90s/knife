package locator

import (
	"gateway/internal/route"
)

type RouteLocator interface {
	GetRoutes() ([]*route.Route, error)
}
