package util

import (
	"log"
	"strings"
)

// HTTP错误日志输出
type httpErrorLogWriter struct{}

func (*httpErrorLogWriter) Write(p []byte) (int, error) {
	message := string(p)
	if !strings.HasPrefix(message, "http: TLS handshake error") &&
		!strings.HasSuffix(message, ": EOF\n") &&
		!strings.HasPrefix(message, "proxy error: context canceled") {
		log.Printf(message)
	}
	return len(p), nil
}

func NewHttpLogger() *log.Logger {
	return log.New(&httpErrorLogWriter{}, "", 0)
}
