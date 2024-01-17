package config

type GateWayConfiguration struct {
	Server *ServerConfiguration
	Logger *LoggerConfiguration
	Routes []*RouteConfiguration
}

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
