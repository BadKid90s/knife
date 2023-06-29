package util

import (
	"os"
)

func ReadConfigFile(configFile string) ([]byte, error) {
	// 读取 YAML 文件内容
	yamlFile, err := os.ReadFile(configFile)
	return yamlFile, err
}
