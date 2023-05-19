package main

import (
	"fmt"
	"gateway/core"
	"gateway/middleware"
	"gateway/network"
	"log"
	"net/http"
)

func logger(ctx *middleware.Context, writer http.ResponseWriter, request *http.Request) error {
	log.Printf("path: %v", request.URL.Path)
	_, err := writer.Write([]byte("hello word"))
	if err != nil {
		return err
	}
	return nil
}

func main() {

	listener, err := network.NewListenTCP(":8080")
	if err != nil {
		log.Printf("app runing err")
	}

	middle := middleware.NewMiddleware()
	err = middle.Register("logger", middleware.HandlerFunc(logger))
	if err != nil {
		log.Printf("register middleware err")
	}

	app := core.NewApp(listener, middle)
	err = app.Start()
	if err != nil {
		log.Printf("app runing err")
	}

	fmt.Println("gateway hello word")
}
