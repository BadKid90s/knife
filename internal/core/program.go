package core

import (
	"fmt"
	"gateway/config/definition"
	"gateway/internal/filter"
	"gateway/internal/filter/global"
	"gateway/internal/handler"
	"gateway/internal/network"
	"gateway/internal/util"
	"gateway/internal/web"
	"gateway/logger"
	"net"
	"net/http"
	"time"
)

func NewApp(configFile string) *ProgramApp {
	startTime := time.Now()

	//打印 banner
	printBanner()

	//解析外部配置
	parseExteriorConfig(configFile)

	logger.Logger.Infof("starting GatewayApplication")

	//分发请求处理器
	dispatcherHandler := web.DispatcherHandlerConstant

	//加载内部配置
	loadInternalConfig(dispatcherHandler)

	return &ProgramApp{
		handler:   dispatcherHandler,
		ip:        definition.GatewayServerDefinition.Ip,
		port:      definition.GatewayServerDefinition.Port,
		startTime: startTime,
	}
}

type ProgramApp struct {
	listener  net.Listener
	handler   *web.DispatcherHandler
	port      int
	ip        string
	startTime time.Time
}

func (a *ProgramApp) Start() {
	//创建监听
	a.createListener()

	//使用核心中间件来服务http
	httpServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           a.handler,
	}
	elapsed := time.Since(a.startTime)
	logger.Logger.Infof("started GatewayApplication in %s", elapsed)

	//监听服务
	err := httpServer.Serve(a.listener)
	if err != nil {
		logger.Logger.Infof("app runing err %s", err)
	}
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

func (a *ProgramApp) createListener() {
	address := fmt.Sprintf("%s:%d", a.ip, a.port)
	listener, err := network.NewListenTCP(address)
	if err != nil {
		logger.Logger.Fatalf("create a listener to send errors, listen to the address: %s ", err)
	}
	logger.Logger.Infof("listener succeeded, listen to the address: %s ", address)
	a.listener = listener
}

func parseExteriorConfig(configFile string) {
	err := definition.ParseConfig(configFile)
	if err != nil {
		logger.Logger.Fatalf("an error occurred in the configuration file parsing [%s] ", err)
	}
}

func loadInternalConfig(dispatcherHandler *web.DispatcherHandler) {
	dispatcherHandler.AddHandler(handler.NewRoutePredicateHandlerMapping())
}

func printBanner() {
	bytes, err := util.ReadConfigFile("config/banner.txt")
	if err != nil {
		logger.Logger.Errorf("loading programe banner err %s", err)
	}
	println(string(bytes))
}
