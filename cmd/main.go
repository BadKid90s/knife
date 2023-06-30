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

	app.Start()

}

func printBanner() {
	bytes, err := util.ReadConfigFile("conf/banner.txt")
	if err != nil {
		log.Fatalf("loading programe banner err %s \n", err)
	}
	println(string(bytes))
}
