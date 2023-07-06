package definition

import (
	"gateway/logger"
	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

func ParseLoggerConfig(buffer []byte) error {
	lcd := &LoggerConfigDefinition{
		Logger: &LoggerDefinition{},
	}
	err := defaults.Set(lcd)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buffer, &lcd)
	if err != nil {
		return err
	}
	logger.NewLogger(lcd.Logger.Level)
	return nil
}
