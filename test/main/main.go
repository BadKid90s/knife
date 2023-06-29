package main

import (
	"fmt"
	"gateway/core"
	"gateway/network"
	"gateway/route"
	"log"
)

func main() {
	var configFile = "conf/application.yaml"
	listener, err := network.NewListenTCP(":8090")
	if err != nil {
		log.Fatalf("app listen err")
	}

	err = route.ParseRouteConfig(configFile)
	if err != nil {
		log.Fatalf("config file parse err")
	}

	app := core.NewApp(listener)
	err = app.Start()
	if err != nil {
		log.Printf("app runing err")
	}

	fmt.Println("gateway hello word")
}
