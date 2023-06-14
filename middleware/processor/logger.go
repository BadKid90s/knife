package processor

import (
	"gateway/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type LoggerMiddleware struct {
	logger *logrus.Logger
}

func (m LoggerMiddleware) Handle(_ *middleware.Context, _ http.ResponseWriter, request *http.Request) (err error) {
	start := time.Now()
	defer func() {
		if err != nil {
			m.logger.Errorf("%s  %v  [%s] 耗时: %s\t   %s  错误: %v", time.Now().Format("2006-01-02 15:04:05.000"), request.RemoteAddr, request.Method, time.Now().Sub(start), request.URL, err)
		} else {
			m.logger.Infof("%s  %v  [%s] 耗时: %s\t   %s", time.Now().Format("2006-01-02 15:04:05.000"), request.RemoteAddr, request.Method, time.Now().Sub(start), request.URL)
		}
	}()
	return nil
}

func init() {
	middleware.RegisteredMiddlewares.RegisterHandler("logger", func(configMap map[string]any) (middleware.Handler, error) {
		return &LoggerMiddleware{
			logger: logrus.New(),
		}, nil
	})
}
