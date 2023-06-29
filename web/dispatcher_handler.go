package web

import "net/http"

var DispatcherHandlerConstant = newDispatcherHandler()

func newDispatcherHandler() *DispatcherHandler {
	return &DispatcherHandler{
		handlerMappings: make([]HandlerMapping, 0),
	}
}

type DispatcherHandler struct {
	handlerMappings []HandlerMapping
}

func (h *DispatcherHandler) AddHandler(mapping HandlerMapping) {
	h.handlerMappings = append(h.handlerMappings, mapping)
}

// Handle 顺序执行中间件
// write: 响应
// request: 请求
func (h *DispatcherHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	exchange := NewServerWebExchange(write, request)
	for _, handlerMapping := range h.handlerMappings {
		handler := handlerMapping.GetHandler(exchange)
		if handler == nil {
			createNotFoundError(exchange)
		} else {
			handler.Handle(exchange)
		}
	}
}

func createNotFoundError(exchange *ServerWebExchange) {
	http.NotFound(exchange.Write, exchange.Request)
}
