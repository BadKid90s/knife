package core

import (
	"gateway/internal/config"
	"gateway/internal/handler"
	"gateway/internal/network"
	"gateway/internal/util"
	"gateway/logger"
	"net"
	"net/http"
	"time"
)

func GatewayApp() *ProgramApp {
	return &ProgramApp{
		configuration: config.DefaultGateWayConfiguration(),
	}
}

type ProgramApp struct {
	configFilePath *string
	configuration  *config.GateWayConfiguration
	listener       net.Listener
}

func (a *ProgramApp) SetConfigFilePath(filepath string) {
	a.configFilePath = &filepath
}

func (a *ProgramApp) Start() {

	startTime := time.Now()

	logger.NewLogger(a.configuration.Logger.Level)

	//打印 banner
	printBanner()

	logger.Logger.Infof("starting GatewayApplication")
	//如果指定了配置文件使用指定的配置
	if a.configFilePath != nil {
		logger.Logger.Infof("loading GateWayConfiguration for %s", *a.configFilePath)
		a.configuration = config.NewGateWayConfiguration(a.configFilePath)
	} else {
		logger.Logger.Infof("use Default GateWayConfiguration")
	}

	//使用核心中间件来服务http
	httpServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           handler.DispatcherHandlerConstant,
	}
	elapsed := time.Since(startTime)
	logger.Logger.Infof("started GatewayApplication in %s", elapsed)

	//创建监听
	a.listener = network.NewListenForIpPort(a.configuration.Server.Ip, a.configuration.Server.Port)

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

func printBanner() {
	bytes, err := util.ReadConfigFile("config/banner.txt")
	if err != nil {
		logger.Logger.Errorf("loading programe banner err %s", err)
	}
	println(string(bytes))
}
