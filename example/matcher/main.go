package main

import (
	"knife"
	"knife/matcher"
	"knife/middleware"
	"log"
	"net/http"
)

func main() {
	//创建http路由
	mux := http.NewServeMux()
	mux.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello 8080 a "))
	})

	//创建中间件链
	chain := knife.NewChain(mux)
	//添加日志记录中间件
	chain.Use(middleware.Logger())
	//添加错误处理中间件
	chain.Use(middleware.Recover())

	//添加响应头是否存在匹配器的自定义中间件
	chain.UseMatcher(matcher.HeaderResponseExists("token"), func(context *knife.Context) {
		log.Printf("token middleware,token:%s ", context.Writer.Header().Get("token"))
		context.Next()
	})

	//添加带有组合匹配器的自定义中间件
	chain.UseMatcher(matcher.Any(matcher.HeaderResponseExists("token"), matcher.HeaderResponseExists("auth")), func(context *knife.Context) {
		log.Printf("token middleware,token:%s ", context.Writer.Header().Get("token"))
		context.Next()
	})

	//启动服务
	_ = http.ListenAndServe(":8080", chain)
}
