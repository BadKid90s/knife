package core

import (
	"gateway/middleware"
	"log"
	"net"
	"net/http"
	"time"
)

func NewApp(listener net.Listener, middleware middleware.Middleware) *ProgramApp {
	return &ProgramApp{
		listener:   listener,
		middleware: middleware,
	}
}

type ProgramApp struct {
	listener   net.Listener
	middleware middleware.Middleware
}

func (a *ProgramApp) Start() error {
	//使用核心中间件来服务http
	httpServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           http.HandlerFunc(a.handleHTTP),
	}
	return httpServer.Serve(a.listener)
}

func (a *ProgramApp) Stop() error {
	if a.listener != nil {
		err := a.listener.Close()
		return err
	}
	return nil
}

func (a *ProgramApp) handleHTTP(write http.ResponseWriter, request *http.Request) {
	err := a.middleware.Handle(write, request)
	if err != nil {
		log.Printf("handle http error %v", err.Error())
	}
}
