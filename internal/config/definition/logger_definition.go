package definition

import (
	"gateway/logger"
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type LoggerConfigDefinition struct {
	Logger *LoggerDefinition `yaml:"logger"`
}

type LoggerDefinition struct {
	Level string `yaml:"level" default:"info"`
}

func ParseLoggerConfig(buffer []byte) (*LoggerConfigDefinition, error) {
	lcd := &LoggerConfigDefinition{
		Logger: &LoggerDefinition{},
	}
	err := defaults.Set(lcd)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(buffer, &lcd)
	if err != nil {
		return nil, err
	}
	logger.NewLogger(lcd.Logger.Level)

	return lcd, nil
}
