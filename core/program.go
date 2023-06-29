package core

import (
	"fmt"
	"gateway/definition"
	"gateway/filter"
	"gateway/filter/global"
	"gateway/handler"
	"gateway/network"
	"gateway/web"
	"log"
	"net"
	"net/http"
	"time"
)

func NewApp(configFile string) *ProgramApp {
	//分发请求处理器
	dispatcherHandler := web.DispatcherHandlerConstant

	//解析外部配置
	parseExteriorConfig(configFile)

	//加载内部配置
	loadInternalConfig(dispatcherHandler)

	return &ProgramApp{
		handler: dispatcherHandler,
	}
}

type ProgramApp struct {
	listener net.Listener
	handler  *web.DispatcherHandler
}

func (a *ProgramApp) Start() error {
	//创建监听
	a.listener = createListener()

	//使用核心中间件来服务http
	httpServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           a.handler,
	}
	return httpServer.Serve(a.listener)
}

func (a *ProgramApp) Stop() error {
	if a.listener != nil {
		err := a.listener.Close()
		return err
	}
	return nil
}

func (a *ProgramApp) Use(filter filter.GatewayFilter) {
	global.Filters = append(global.Filters, filter)
}

func parseExteriorConfig(configFile string) {
	err := definition.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("an error occurred in the configuration file parsing [%s] \n", err)
	}
}

func loadInternalConfig(dispatcherHandler *web.DispatcherHandler) {
	dispatcherHandler.AddHandler(handler.NewRoutePredicateHandlerMapping())
}

func createListener() net.Listener {
	address := fmt.Sprintf("%s:%d", definition.GatewayServerDefinition.Ip, definition.GatewayServerDefinition.Port)
	listener, err := network.NewListenTCP(address)
	if err != nil {
		log.Fatalf("create a listener to send errors, listen to the address [%s] \n", err)
	}
	log.Printf("listener succeeded ,listen to the address [%s] \n", address)
	return listener
}
