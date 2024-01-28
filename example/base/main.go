package main

import (
	"knife"
	"knife/middleware"
	"log"
	"net/http"
	"time"
)

func main() {

	//Create a middleware chain
	//first plan
	chain := planOne()

	//second plan
	//chain := planTwo()

	//the third plan
	//chain := planThree()

	//start serve
	_ = http.ListenAndServe(":8080", chain)
}

// plan one
func planOne() *knife.Chain {
	//create  middleware chain
	chain := knife.NewChain()
	//add logger middleware
	chain.Use(middleware.Logger())
	//add recover middleware
	chain.Use(middleware.Recover())

	//add custom implemented middleware
	chain.Use(func(context *knife.Context) {
		start := time.Now()
		defer func() {
			duration := time.Now().Sub(start)
			log.Printf("process all middleware,time consumption %s ", duration)
		}()
		context.Next()
	})
	return chain
}

// plan two
// add middleware to the constructor
func planTwo() *knife.Chain {
	//create a middleware chain and add logging and error handling middleware
	chain := knife.NewChain(middleware.Logger(), middleware.Recover())

	//add custom implemented middleware
	chain.Use(func(context *knife.Context) {
		start := time.Now()
		defer func() {
			duration := time.Now().Sub(start)
			log.Printf("process all middleware,time consumption %s ", duration)
		}()
		context.Next()
	})
	return chain
}

// plan three
// use chain call
func planThree() *knife.Chain {
	//create a middleware chain and use the chain to add logging, error handling, and custom middleware
	chain := knife.NewChain().
		Use(middleware.Logger()).
		Use(middleware.Recover()).
		Use(func(context *knife.Context) {
			start := time.Now()
			defer func() {
				duration := time.Now().Sub(start)
				log.Printf("process all middleware,time consumption %s ", duration)
			}()
			context.Next()
		})

	return chain
}
