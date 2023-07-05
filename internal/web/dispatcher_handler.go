package web

import (
	"gateway/logger"
	"log"
	"net/http"
	"time"
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
	startTime := time.Now()
	logger.Logger.Infof("dispatcher handler received request. url: %s ", request.URL)

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

	elapsed := time.Since(startTime)
	logger.Logger.Infof("dispatcher handler process sucess in %s", elapsed)
}

func createNotFoundError(exchange *ServerWebExchange) {
	http.NotFound(exchange.Write, exchange.Request)
}
