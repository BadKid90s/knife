package processor

import (
	"gateway/middleware"
	"gateway/util"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxyConfiguration struct {
	Target string `config:"target"`
}

type ReverseProxyMiddleware struct {
	proxy *httputil.ReverseProxy
}

func (r *ReverseProxyMiddleware) Handle(_ *middleware.Context, writer http.ResponseWriter, request *http.Request) (bool, error) {
	r.proxy.ServeHTTP(writer, request)
	return true, nil
}

func init() {
	middleware.RegisteredMiddlewares.RegisterHandler("proxy", func(configMap map[string]any) (middleware.Handler, error) {
		//解析配置
		config := &ReverseProxyConfiguration{}
		err := util.UnpackConfig(configMap, config)
		if err != nil {
			return nil, err
		}
		//构造反向代理
		targetURL, err := url.Parse(config.Target)
		if err != nil {
			return nil, err
		}
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ErrorLog = util.NewHttpLogger()
		return &ReverseProxyMiddleware{
			proxy: proxy,
		}, err
	})

}
