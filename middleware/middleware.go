package middleware

import (
	"errors"
	"net/http"
)

func NewMiddleware() Middleware {
	return Middleware{
		ctx:                  &Context{},
		registeredMiddleware: make(map[string]MiddlewareHandler),
	}

}

type Middleware struct {
	ctx                  *Context
	registeredMiddleware map[string]MiddlewareHandler
}

func (m *Middleware) Handle(write http.ResponseWriter, request *http.Request) error {

	for _, value := range m.registeredMiddleware {
		err := value.Handle(m.ctx, write, request)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Middleware) register(id string, handler MiddlewareHandler) error {
	_, exist := m.registeredMiddleware[id]
	if exist == false {
		m.registeredMiddleware[id] = handler
		return nil
	}
	return RepeatRegisterErr
}

// RepeatRegisterErr 中间件重复注册
var RepeatRegisterErr = errors.New("middleware Repeat Register")
