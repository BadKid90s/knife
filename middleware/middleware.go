package middleware

import (
	"errors"
	"net/http"
)

// RegisteredMiddlewares 已经注册的中间件
var RegisteredMiddlewares = newMiddleware()

// NewMiddleware 中间件的构造函数
func newMiddleware() Middleware {
	return Middleware{
		registeredHandlerConstructor: make(map[string]HandlerConstructor, 0),
		handlers:                     make([]Handler, 0),
	}

}

// Middleware 中间件
type Middleware struct {
	//注册的中间件处理器构造函数
	registeredHandlerConstructor map[string]HandlerConstructor
	//有序的中间件
	handlers []Handler
}

// Handle 顺序执行中间件
// write: 响应
// request: 请求
func (m *Middleware) Handle(write http.ResponseWriter, request *http.Request) error {
	createContext(m.handlers).Next(write, request)
	return nil
}

// RegisterHandler 注册中间件(接口实现)
// id: 中间件的key
// constructor: 中间件构造器
func (m *Middleware) RegisterHandler(id string, constructor HandlerConstructor) {
	_, exist := m.registeredHandlerConstructor[id]
	if exist == false {
		m.registeredHandlerConstructor[id] = constructor
	}
}

func (m *Middleware) BuildHandler(id string, config map[string]any) error {
	if constructor, exist := m.registeredHandlerConstructor[id]; exist {
		handler, err := constructor(config)
		if err != nil {
			return err
		}
		m.handlers = append(m.handlers, handler)
		return nil
	}
	return NotFoundRegisterErr
}

// HandlerConstructor 中间件处理器构造函数
type HandlerConstructor func(configMap map[string]any) (Handler, error)

// HandlerConstructorFunc 中间件处理器构造函数
type HandlerConstructorFunc func(configMap map[string]any) (HandlerFunc, error)

// NotFoundRegisterErr 中间件没有找到
var NotFoundRegisterErr = errors.New("middleware Repeat Register")
