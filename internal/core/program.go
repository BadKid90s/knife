package core

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/handler"
	"gateway/internal/network"
	"gateway/logger"
	"net"
	"net/http"
	"time"
)

// GatewayApp 创建程序
func GatewayApp() *ProgramApp {
	return &ProgramApp{}
}

type ProgramApp struct {
	configFilePath *string
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

	//打印 banner
	printBanner()

	config.NewGatewayConfiguration(a.configFilePath)

	//创建监听
	ip := config.GlobalGatewayConfiguration.Server.Ip
	port := config.GlobalGatewayConfiguration.Server.Port
	a.listener = network.NewListenForIpPort(ip, port)

	elapsed := time.Since(startTime)
	logger.Logger.TagLogger("core").Infof("started gatewayApplication in %s", elapsed)

	run(a.listener)
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
	fmt.Println("____       _                           ")
	fmt.Println("/ ___| __ _| |_ _____      ____ _ _   _")
	fmt.Println("| |  _ / _` | __/ _ \\ \\ / / _` | | | |")
	fmt.Println("| |_| | (_| | ||  __/\\ V  V / (_| | |_| |")
	fmt.Println("\\____|\\__,_|\\__\\___| \\_/\\_/ \\__,_|\\__, |")
	fmt.Println("                                    |___/")
}

func run(listener net.Listener) {
	//使用核心中间件来服务http
	httpServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           handler.NewDispatcherHandler(),
	}
	//监听服务
	err := httpServer.Serve(listener)

	if err != nil {
		logger.Logger.TagLogger("core").Errorf("app runing err %s", err)
	}
}
