package main

import (
	"gateway/internal/core"
)

func main() {

	var configFile = "config/application.yaml"

	app := core.NewApp(configFile)

	app.Start()

}
