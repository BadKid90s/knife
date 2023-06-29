package core

import (
	"gateway/definition"
	"gateway/handler"
	"gateway/web"
	"log"
	"net"
	"net/http"
	"time"
)

func NewApp(listener net.Listener, configFile string) *ProgramApp {
	return &ProgramApp{
		listener:   listener,
		handler:    web.DispatcherHandlerConstant,
		configFile: configFile,
	}
}

type ProgramApp struct {
	listener   net.Listener
	handler    *web.DispatcherHandler
	configFile string
}

func (a *ProgramApp) Start() error {
	//解析外部配置
	a.parseExteriorConfig()

	//加载内部配置
	a.loadInternalConfig()

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

func (a *ProgramApp) parseExteriorConfig() {
	err := definition.ParseRouteConfig(a.configFile)
	if err != nil {
		log.Fatalf("config file parse err")
	}
}

func (a *ProgramApp) loadInternalConfig() {
	a.handler.AddHandler(handler.NewRoutePredicateHandlerMapping())
}
