package network

import (
	"net"
	"time"
)

// NewListenTCP 建立TCP连接
func NewListenTCP(address string) (net.Listener, error) {
	underlying, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &TCPListener{
		Listener: underlying,
		denyTTL:  5 * time.Minute,
	}, nil
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
