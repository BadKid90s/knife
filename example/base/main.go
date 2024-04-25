package main

import (
	"fmt"
	"knife"
	"knife/middleware/cache"
	"knife/middleware/gzip"
	"knife/middleware/logger"
	"knife/middleware/proxy"
	"knife/middleware/recover"
	"log"
	"net/http"
	"time"
)

func main() {

	//Create a middleware chain
	//first plan
	//chain := planOne()

	//second plan
	//chain := planTwo()

	//the third plan
	//chain := planThree()

	//gzip middleware
	//chain := gzipFun()

	//cache middleware
	//chain := cacheFun()

	//cache middleware
	chain := loadBalance()

	//start serve
	err := http.ListenAndServe(":8080", chain)
	if err != nil {
		panic(err)
	}
}

// plan one
func planOne() *knife.Chain {
	//create  middleware chain
	chain := knife.NewChain()
	//add logger middleware
	chain.Use(logger.Logger())
	//add recover middleware
	chain.Use(recover.Recover())

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
	chain := knife.NewChain(logger.Logger(), recover.Recover())

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
		Use(logger.Logger()).
		Use(recover.Recover()).
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

func gzipFun() *knife.Chain {
	chain := knife.NewChain().
		Use(logger.Logger()).
		Use(recover.Recover()).
		//Use(gzip.Default()).
		Use(gzip.Gzip(1024)).
		Use(func(context *knife.Context) {
			data := "Gzip是一种压缩文件格式并且也是一个在类 Unix 上的一种文件解压缩的软件，通常指GNU计划的实现，此处的gzip代表GNU zip。" +
				"也经常用来表示gzip这种文件格式。软件的作者是Jean-loup Gailly和Mark Adler。在1992年10月31日第一次公开发布，版本号0.1，1993年2月，发布了1.0版本。" +
				"Gzip是一种压缩文件格式并且也是一个在类 Unix 上的一种文件解压缩的软件，通常指GNU计划的实现，此处的gzip代表GNU zip。" +
				"也经常用来表示gzip这种文件格式。软件的作者是Jean-loup Gailly和Mark Adler。在1992年10月31日第一次公开发布，版本号0.1，1993年2月，发布了1.0版本。" +
				"Gzip是一种压缩文件格式并且也是一个在类 Unix 上的一种文件解压缩的软件，通常指GNU计划的实现，此处的gzip代表GNU zip。" +
				"也经常用来表示gzip这种文件格式。软件的作者是Jean-loup Gailly和Mark Adler。在1992年10月31日第一次公开发布，版本号0.1，1993年2月，发布了1.0版本。"
			_, err := context.Writer.Write([]byte(data))
			if err != nil {
				panic(fmt.Sprintf("writer data error %s", err))
			}
			context.Writer.WriteHeader(http.StatusOK)
		})
	return chain
}

func cacheFun() *knife.Chain {
	chain := knife.NewChain().
		Use(logger.Logger()).
		Use(recover.Recover()).
		Use(cache.Cache(30, 60)).
		Use(func(context *knife.Context) {
			data := "Hello World"
			_, err := context.Writer.Write([]byte(data))
			if err != nil {
				panic(fmt.Sprintf("writer data error %s", err))
			}
			context.Writer.WriteHeader(http.StatusOK)
		})
	return chain
}

func loadBalance() *knife.Chain {
	nodes := []*proxy.ServiceNode{
		{
			Address: "127.0.0.1:8080",
			Weight:  1,
		},
		{
			Address: "127.0.0.2:8080",
			Weight:  1,
		},
		{
			Address: "127.0.0.3:8080",
			Weight:  1,
		},
	}
	chain := knife.NewChain().
		Use(logger.Logger()).
		Use(recover.Recover()).
		Use(proxy.LoadBalanceProxy(proxy.LoadBalanceRandom, nodes)).
		Use(func(context *knife.Context) {
			data := "Hello World"
			_, err := context.Writer.Write([]byte(data))
			if err != nil {
				panic(fmt.Sprintf("writer data error %s", err))
			}
			context.Abort(http.StatusOK)
		})
	return chain
}
