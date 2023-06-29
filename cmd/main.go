package main

import (
	"gateway/core"
	"log"
)

func main() {
	var configFile = "conf/application.yaml"

	app := core.NewApp(configFile)

	err := app.Start()
	if err != nil {
		log.Printf("app runing err \n")
	}
}
