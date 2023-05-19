package middleware

import (
	"errors"
	"net/http"
)

func NewMiddleware() Middleware {
	return Middleware{
		ctx:                  &Context{},
		registeredMiddleware: make(map[string]Handler),
	}

}

type Middleware struct {
	ctx                  *Context
	registeredMiddleware map[string]Handler
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

func (m *Middleware) Register(id string, handler Handler) error {
	_, exist := m.registeredMiddleware[id]
	if exist == false {
		m.registeredMiddleware[id] = handler
		return nil
	}
	return RepeatRegisterErr
}

// RepeatRegisterErr 中间件重复注册
var RepeatRegisterErr = errors.New("middleware Repeat Register")
