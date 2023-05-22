package middleware

import (
	"errors"
	"net/http"
)

// NewMiddleware 中间件的构造函数
func NewMiddleware() Middleware {
	return Middleware{
		ctx:                  &Context{},
		registeredMiddleware: make(map[string]Handler),
		middlewareKeys:       make([]string, 0),
	}

}

// Middleware 中间件
type Middleware struct {
	//上下文环境
	ctx *Context
	//注册的中间件
	registeredMiddleware map[string]Handler
	//有序的中间件key
	middlewareKeys []string
}

// Handle 按照`middlewareKeys`中的顺序执行中间件
// write: 响应
// request: 请求
func (m *Middleware) Handle(write http.ResponseWriter, request *http.Request) error {

	for _, value := range m.middlewareKeys {
		err := m.registeredMiddleware[value].Handle(m.ctx, write, request)
		if err != nil {
			return err
		}
	}
	return nil
}

// RegisterHandler 注册中间件(接口实现)
// id: 中间件的key
// handler: 中间件
func (m *Middleware) RegisterHandler(id string, handler Handler) error {
	_, exist := m.registeredMiddleware[id]
	if exist == false {
		m.registeredMiddleware[id] = handler
		m.middlewareKeys = append(m.middlewareKeys, id)
		return nil
	}
	return RepeatRegisterErr
}

// RegisterHandlerFunc 注册中间件（方法实现）
// id: 中间件的key
// handler: 中间件
func (m *Middleware) RegisterHandlerFunc(id string, handlerFunc HandlerFunc) error {
	return m.RegisterHandler(id, handlerFunc)
}

// RepeatRegisterErr 中间件重复注册
var RepeatRegisterErr = errors.New("middleware Repeat Register")
