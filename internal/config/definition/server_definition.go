package definition

import (
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type ServerConfigDefinition struct {
	Server *ServerDefinition `yaml:"server"`
}
type ServerDefinition struct {
	Ip   string `yaml:"ip" default:"0.0.0.0"`
	Port int    `yaml:"port" default:"8080"`
}

func ParseServerConfig(buffer []byte) (*ServerConfigDefinition, error) {
	gsd := &ServerConfigDefinition{
		Server: &ServerDefinition{},
	}
	err := defaults.Set(gsd)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buffer, &gsd)
	if err != nil {
		return nil, err
	}
	return gsd, nil
}
