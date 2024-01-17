package handler

import (
	"gateway/internal/web"
	"gateway/logger"
	"net/http"
	"time"
)

var DispatcherHandlerConstant = newDispatcherHandler()

func newDispatcherHandler() *DispatcherHandler {
	return &DispatcherHandler{
		handlerMappings: []HandlerMapping{NewRoutePredicateHandlerMapping()},
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

	exchange := web.NewServerWebExchange(write, request)
	for _, handlerMapping := range h.handlerMappings {
		handler, err := handlerMapping.GetHandler(exchange)
		if err != nil {
			logger.Logger.Errorf(err.Error())
			createOtherError(exchange, err.Error())
		} else if handler == nil {
			createNotFoundError(exchange)
		} else {
			handler.Handle(exchange)
		}
	}

	elapsed := time.Since(startTime)
	logger.Logger.Infof("dispatcher handler process sucess in %s", elapsed)
}

func createNotFoundError(exchange *web.ServerWebExchange) {
	http.NotFound(exchange.Write, exchange.Request)
}
func createOtherError(exchange *web.ServerWebExchange, msg string) {
	http.Error(exchange.Write, msg, 500)
}
