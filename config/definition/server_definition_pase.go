package definition

import (
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

var GatewayServerDefinition *ServerDefinition

func ParseServerConfig(buffer []byte) error {
	gsd := &ServerConfigDefinition{
		Server: &ServerDefinition{},
	}
	err := defaults.Set(gsd)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buffer, &gsd)
	if err != nil {
		return err
	}
	GatewayServerDefinition = gsd.Server
	return nil
}
