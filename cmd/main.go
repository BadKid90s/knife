package main

import (
	"gateway/internal/core"
)

func main() {

	app := core.GatewayApp()
	//app.SetConfigFilePath("config/application.yml")
	app.Start()

}
