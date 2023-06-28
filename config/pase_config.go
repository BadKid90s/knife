package config

import (
	"gateway/route"
	"gateway/util"
	"gopkg.in/yaml.v2"
	"os"
)

var configFile = "conf/application.yaml"
var RouterDefinitions = make([]*route.Definition, 0)

type GatewayRoutesDefinition struct {
	Routes []*route.Definition `config:"routes"`
}

func ParsePredicateConfig() error {
	configMap, err := readConfigFile()
	if err != nil {
		return err
	}
	//解析配置
	config := &GatewayRoutesDefinition{
		Routes: make([]*route.Definition, 0),
	}
	err = util.UnpackConfig(configMap, config)
	if err != nil {
		return err
	}
	RouterDefinitions = config.Routes
	return nil
}

func readConfigFile() (map[string]any, error) {
	// 读取 YAML 文件内容
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	// 定义一个 map 类型的变量
	configMap := make(map[string]any)

	// 将 YAML 文件内容解析为 map 类型
	err = yaml.Unmarshal(yamlFile, &configMap)
	return configMap, err
}
