package main

import (
	"knife"
	"knife/matcher/combination"
	"knife/matcher/header"
	"knife/matcher/path"
	"knife/middleware/logger"
	"knife/middleware/recover"
	"log"
	"net/http"
)

func main() {

	//创建中间件链
	chain := knife.NewChain()
	//添加日志记录中间件
	chain.Use(logger.Logger())
	//添加错误处理中间件
	chain.Use(recover.Recover())

	//添加响应头是否存在匹配器的自定义中间件
	chain.UseMatcher(header.HeaderResponseExists("token"), func(context *knife.Context) {
		log.Printf("token middleware,token:%s ", context.Writer.Header().Get("token"))
		context.Next()
	})

	//添加带有组合匹配器的自定义中间件
	chain.UseMatcher(combination.Any(header.HeaderResponseExists("token"), header.HeaderResponseExists("auth")), func(context *knife.Context) {
		log.Printf("token middleware,token:%s ", context.Writer.Header().Get("token"))
		context.Next()
	})

	//添加带路径匹配器的自定义中间件
	chain.UseMatcher(path.PathPrefix("/hello"), func(context *knife.Context) {
		log.Println("pathPrefix matcher")
		context.Next()
	})

	//启动服务
	_ = http.ListenAndServe(":8080", chain)
}
