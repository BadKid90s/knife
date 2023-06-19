package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 创建一个 ServeMux 实例
	mux := http.NewServeMux()

	// 为 /hello/* 配置处理函数
	mux.HandleFunc("/hello/a", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello, A!")
		if err != nil {
			log.Printf("println err")
		}
	})
	mux.Handle("/hello/b", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello, B!")
		if err != nil {
			log.Printf("println err")
		}
	}))

	// 为 /goodbye/* 配置处理函数
	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Goodbye, world!")
		if err != nil {
			log.Printf("println err")
		}
	}
	mux.Handle("/goodbye/a", http.HandlerFunc(goodbyeHandler))

	// 监听网络地址，并提供 HTTP 服务
	err := http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		fmt.Println("ListenAndServe error:", err)
	}
}
