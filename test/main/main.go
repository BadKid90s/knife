package main

import (
	"fmt"
	"gateway/core"
	"gateway/middleware"
	"gateway/network"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func main() {

	listener, err := network.NewListenTCP(":8090")
	if err != nil {
		log.Fatalf("app listen err")
	}

	// 读取 YAML 文件内容
	yamlFile, err := os.ReadFile("configs/application.yaml")
	if err != nil {
		panic(err)
	}

	// 定义一个 map 类型的变量
	configMap := make(map[string]any)

	// 将 YAML 文件内容解析为 map 类型
	err = yaml.Unmarshal(yamlFile, &configMap)
	if err != nil {
		panic(err)
	}

	middle := middleware.RegisteredMiddlewares
	err = middle.BuildHandler("logger", configMap)
	if err != nil {
		log.Printf("buildHandler err")
	}
	err = middle.BuildHandler("proxy", configMap)
	if err != nil {
		log.Printf("buildHandler err")
	}
	app := core.NewApp(listener, middle)
	err = app.Start()
	if err != nil {
		log.Printf("app runing err")
	}

	fmt.Println("gateway hello word")
}
