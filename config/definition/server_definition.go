package definition

type ServerConfigDefinition struct {
	Server *ServerDefinition `yaml:"server"`
}
type ServerDefinition struct {
	Ip   string `yaml:"ip" default:"0.0.0.0"`
	Port int    `yaml:"port" default:"8080"`
}
