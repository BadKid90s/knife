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

// GatewayApp 创建程序
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

// SetConfigFilePath 设置配置文件路径
// 如果设置了就使用自定义的配置否则使用默认的配置
// 默认的配置： config.DefaultGateWayConfiguration
func (a *ProgramApp) SetConfigFilePath(filepath string) {
	a.configFilePath = &filepath
}

// Start 启动程序
func (a *ProgramApp) Start() {

	startTime := time.Now()

	logger.NewLogger(a.configuration.Logger.Level)

	//打印 banner
	printBanner()

	logger.Logger.Infof("starting gatewayApplication")
	//如果指定了配置文件使用指定的配置
	if a.configFilePath != nil {
		logger.Logger.Infof("loading gateWayConfiguration for %s", *a.configFilePath)
		a.configuration = config.NewGateWayConfiguration(a.configFilePath)
	} else {
		logger.Logger.Infof("use default gateWayConfiguration")
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
	logger.Logger.Infof("started gatewayApplication in %s", elapsed)

	//创建监听
	a.listener = network.NewListenForIpPort(a.configuration.Server.Ip, a.configuration.Server.Port)

	//监听服务
	err := httpServer.Serve(a.listener)
	if err != nil {
		logger.Logger.Infof("app runing err %s", err)
	}
}

// Stop 停止程序
func (a *ProgramApp) Stop() error {
	if a.listener != nil {
		err := a.listener.Close()
		return err
	}
	return nil
}

// 打印banner
func printBanner() {
	bytes, err := util.ReadConfigFile("config/banner.txt")
	if err != nil {
		logger.Logger.Errorf("loading programe banner err %s", err)
	}
	println(string(bytes))
}
