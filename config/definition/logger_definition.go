package definition

type LoggerConfigDefinition struct {
	Logger *LoggerDefinition `yaml:"logger"`
}
type LoggerDefinition struct {
	Level string `yaml:"level" default:"info"`
}
