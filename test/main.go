package main

import (
	"fmt"
	"gateway/core"
	"gateway/middleware"
	"gateway/network"
	"log"
)

func main() {

	listener, err := network.NewListenTCP(":8090")
	if err != nil {
		log.Fatalf("app listen err")
	}

	middle := middleware.RegisteredMiddlewares
	err = middle.BuildHandler("logger", make(map[string]any))
	if err != nil {
		log.Printf("buildHandler err")
	}
	configMap := make(map[string]any)
	configMap["target"] = "http://localhost:5173"
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
