package core

import (
	"gateway/web"
	"net"
	"net/http"
	"time"
)

func NewApp(listener net.Listener) *ProgramApp {
	return &ProgramApp{
		listener: listener,
		handler:  web.DispatcherHandlerConstant,
	}
}

type ProgramApp struct {
	listener net.Listener
	handler  *web.DispatcherHandler
}

func (a *ProgramApp) Start() error {
	//使用核心中间件来服务http
	httpServer := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       5 * time.Minute,
		Handler:           a.handler,
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
