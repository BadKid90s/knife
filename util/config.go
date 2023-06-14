package util

import (
	"errors"
	"github.com/elastic/go-ucfg"
)

// UnpackConfig 解包配置
func UnpackConfig(config map[string]any, to any) error {
	if config == nil {
		return errors.New("config map cannot be null")
	}
	configInstance, err := ucfg.NewFrom(config)
	if err != nil {
		return err
	}
	return configInstance.Unpack(to)
}
