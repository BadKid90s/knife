package main

import (
	"gateway/internal/core"
	"gateway/internal/util"
	"log"
)

func main() {
	printBanner()

	var configFile = "config/application.yaml"

	app := core.NewApp(configFile)

	app.Start()

}

func printBanner() {
	bytes, err := util.ReadConfigFile("config/banner.txt")
	if err != nil {
		log.Fatalf("loading programe banner err %s \n", err)
	}
	println(string(bytes))
}
