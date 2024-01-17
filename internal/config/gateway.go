package config

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
