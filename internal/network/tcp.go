package network

import (
	"fmt"
	"gateway/logger"
	"net"
	"time"
)

// NewListenForAddress 建立TCP连接
func NewListenForAddress(address string) (net.Listener, error) {
	underlying, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &TCPListener{
		Listener: underlying,
		denyTTL:  5 * time.Minute,
	}, nil
}

func NewListenForIpPort(ip string, port int) net.Listener {
	address := fmt.Sprintf("%s:%d", ip, port)
	listener, err := NewListenForAddress(address)
	if err != nil {
		logger.Logger.TagLogger("network").Fatalf("create a listener to send errors, listen to the address: %s ", err)
	}
	logger.Logger.TagLogger("network").Infof("listener succeeded, listen to the address: %s ", address)
	return listener
}

type TCPListener struct {
	net.Listener
	denyTTL time.Duration
}

func (l *TCPListener) Accept() (net.Conn, error) {
	for {
		//获取底层连接
		conn, err := l.Listener.Accept()
		if err != nil {
			return nil, err
		}
		//正常连接
		return conn, nil
	}
}

func (l *TCPListener) Close() error {
	return l.Listener.Close()
}
