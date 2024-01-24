package main

import (
	"knife"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello a "))
	})

	chain := knife.NewChain(mux).Use(func(context *knife.Context) {
		start := time.Now()
		defer func() {
			duration := time.Now().Sub(start)
			log.Printf("process all middleware,time consumption %s ", duration)
		}()
		context.Next()
	}).Use(func(context *knife.Context) {
		log.Printf("logger middleware ,path:%s ", context.Req.URL.Path)
		context.Next()
	}).UseMatcher(func(c *knife.Context) bool {
		return true
	}, func(context *knife.Context) {

	})

	_ = http.ListenAndServe(":8080", chain)
}
