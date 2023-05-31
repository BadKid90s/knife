package processor

import (
	"gateway/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type LoggerMiddleware struct {
	logger *logrus.Entry
}

func (m LoggerMiddleware) Handle(_ *middleware.Context, _ http.ResponseWriter, request *http.Request) (err error) {
	start := time.Now()
	defer func() {
		if err != nil {
			m.logger.Errorf("[ERR] %s %s 耗时: %v 来自: %s 错误: %v", request.Method, request.URL, time.Now().Sub(start), request.RemoteAddr, err)
		} else {
			m.logger.Infof("[%d] %s 耗时: %v 来自: %s", request.Method, request.URL, time.Now().Sub(start), request.RemoteAddr)
		}
	}()
	return nil
}

func init() {

}
