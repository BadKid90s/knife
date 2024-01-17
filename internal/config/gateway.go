package config

import (
	"errors"
	"gateway/internal/config/definition"
	"gateway/internal/util"
	"gateway/logger"
)

// GateWayConfiguration 程序配置
type GateWayConfiguration struct {
	//服务
	Server *ServerConfiguration
	//日志
	Logger *LoggerConfiguration
	//路由
	Routes []*RouteConfiguration
}

// DefaultGateWayConfiguration 默认的程序配置
// 服务端口：8080
// 日志级别：info
// 路由：无
func DefaultGateWayConfiguration() *GateWayConfiguration {
	return &GateWayConfiguration{
		Server: &ServerConfiguration{
			Ip:   "0.0.0.0",
			Port: 8080,
		},
		Logger: &LoggerConfiguration{
			Level: "info",
		},
		Routes: make([]*RouteConfiguration, 0),
	}
}

func NewGateWayConfiguration(filepath *string) *GateWayConfiguration {
	configuration := DefaultGateWayConfiguration()
	err := parseConfig(filepath, configuration)
	if err != nil {
		logger.Logger.Fatalf("parse config failed，err:%s ", err)
	}
	return configuration
}

func parseConfig(configFile *string, configuration *GateWayConfiguration) error {

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
	configuration.Routes = routes

	return nil

}
