package main

import (
	"fmt"
	"gateway/config"
	"gateway/core"
	"gateway/network"
	"log"
)

func main() {

	listener, err := network.NewListenTCP(":8090")
	if err != nil {
		log.Fatalf("app listen err")
	}

	err = config.ParsePredicateConfig()
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
