package main

import (
	"gateway/internal/core"
)

func main() {

	app := core.GatewayApp()
	app.SetConfigFilePath("config/application.yaml")
	app.Start()

}
