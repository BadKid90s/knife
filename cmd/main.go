package main

import (
	"gateway/core"
	"gateway/util"
	"log"
)

func main() {
	printBanner()

	var configFile = "conf/application.yaml"

	app := core.NewApp(configFile)

	err := app.Start()
	if err != nil {
		log.Printf("app runing err %s \n", err)
	}
}

func printBanner() {
	bytes, err := util.ReadConfigFile("conf/banner.txt")
	if err != nil {
		log.Fatalf("loading programe banner err %s \n", err)
	}
	println(string(bytes))
}
