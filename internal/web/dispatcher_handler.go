package web

import (
	"log"
	"net/http"
)

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
	log.Printf("dispatcher handler received request. url [%s] \n", request.URL)
	exchange := NewServerWebExchange(write, request)
	for _, handlerMapping := range h.handlerMappings {
		handler, err := handlerMapping.GetHandler(exchange)
		if err != nil {
			log.Printf(err.Error())
		} else if handler == nil {
			createNotFoundError(exchange)
		} else {
			handler.Handle(exchange)
		}
	}
}

func createNotFoundError(exchange *ServerWebExchange) {
	http.NotFound(exchange.Write, exchange.Request)
}
