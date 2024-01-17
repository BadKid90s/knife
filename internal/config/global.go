package config

import (
	"errors"
	"gateway/internal/config/definition"
	"gateway/internal/util"
	"gateway/logger"
)

var GlobalGatewayConfiguration *GatewayConfiguration

// GatewayConfiguration 程序配置
type GatewayConfiguration struct {
	//服务
	Server *ServerConfiguration
	//日志
	Logger *LoggerConfiguration
	//路由
	Router *GatewayRoutesConfiguration
}

// DefaultGateWayConfiguration 默认的程序配置
// 服务端口：8080
// 日志级别：info
// 路由：无
func defaultGateWayConfiguration() *GatewayConfiguration {
	return &GatewayConfiguration{
		Server: &ServerConfiguration{
			Ip:   "0.0.0.0",
			Port: 8080,
		},
		Logger: &LoggerConfiguration{
			Level: "info",
		},
		Router: &GatewayRoutesConfiguration{
			Routes: make([]*RouteConfiguration, 0),
		},
	}
}

func NewGatewayConfiguration(filepath *string) {

	configuration := defaultGateWayConfiguration()

	//如果指定了配置文件使用指定的配置
	if filepath != nil {
		err := parseConfig(filepath, configuration)
		logger.NewLogger(configuration.Logger.Level)
		logger.Logger.Infof("loading gateWayConfiguration for %s", *filepath)
		if err != nil {
			logger.Logger.Fatalf("parse config failed，err:%s ", err)
		}
	} else {
		logger.NewLogger(configuration.Logger.Level)
		logger.Logger.Infof("use default gateWayConfiguration")
	}
	GlobalGatewayConfiguration = configuration
}

func parseConfig(configFile *string, configuration *GatewayConfiguration) error {

	buffer, err := util.ReadConfigFile(*configFile)
	if err != nil {
		return err
	}

	//解析日志
	loggerConfig, err := definition.ParseLoggerConfig(buffer)
	if err != nil {
		return errors.New("an error occurred while parsing the logger configuration")
	}
	configuration.Logger.Level = loggerConfig.Logger.Level

	//解析服务配置
	serverConfig, err := definition.ParseServerConfig(buffer)
	if err != nil {
		return errors.New("an error occurred while parsing the server configuration")
	}
	configuration.Server.Ip = serverConfig.Server.Ip
	configuration.Server.Port = serverConfig.Server.Port

	routeConfig, err := definition.ParseRouteConfig(buffer)
	if err != nil {
		return errors.New("an error occurred while parsing the routes configuration")
	}

	routes := make([]*RouteConfiguration, len(routeConfig.Routes))
	for i, v := range routeConfig.Routes {

		var predicates = make([]*PredicateConfiguration, len(v.PredicateDefinitions))
		for j, item := range v.PredicateDefinitions {
			predicates[j] = &PredicateConfiguration{
				Name: item.Name,
				Args: item.Args,
			}
		}

		routes[i] = &RouteConfiguration{
			Id:                     v.Id,
			Uri:                    v.Uri,
			Order:                  v.Order,
			PredicateConfiguration: predicates,
		}
	}
	configuration.Router.Routes = routes

	return nil

}
