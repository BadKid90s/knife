package main

import (
	"fmt"
	"gateway/core"
	"gateway/network"
	"log"
)

func main() {
	var configFile = "conf/application.yaml"
	listener, err := network.NewListenTCP(":8090")
	if err != nil {
		log.Fatalf("app listen err")
	}

	app := core.NewApp(listener, configFile)
	err = app.Start()
	if err != nil {
		log.Printf("app runing err")
	}

	fmt.Println("gateway hello word")
}
