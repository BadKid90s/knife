# Knife
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/BadKid90s/knife?sort=semver)](https://github.com/BadKid90s/knife/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/minekube/gate?logo=go)](https://golang.org/doc/devel/release.html)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/BadKid90s/knife/go.yml?branch=main&style=flat-square)](https://github.com/BadKid90s/knife/blob/main/.github/workflows/go.yml)
[![GitHub issues](https://img.shields.io/github/issues/BadKid90s/knife?style=flat-square)](https://github.com/BadKid90s/knife/issues)

Knife Is The Lightweight Middleware For Golang

- Low memory usage
- High performance
- Developed in Go 
- Embrace the cloud native era!


# Getting Started

Gate is designed with developers in mind.

All you need to get started is a working Go environment. 
You can find the Go installation instructions [here](https://go.dev/doc/install).

Once you have Go installed, you create a new Go module and add Knife as a dependency:
```shell
mkdir knife-demo; 
cd knife-demo
go mod init knife-demo
go get github.com/BadKid90s/knife
```

Add and initialize your program and execute it, that's it!
```go
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

```


# Learning by Example

The best way to learn how to extend Knife is by looking at some examples.
If you want to see a complete Go project that uses Knife, [check out the Simple example](https://github.com/BadKid90s/knife/tree/main/example).
