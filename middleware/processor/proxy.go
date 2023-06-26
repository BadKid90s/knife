package processor

import (
	"fmt"
	"gateway/matcher/host"
	"gateway/matcher/method"
	"gateway/matcher/path"
	timeMatch "gateway/matcher/time"
	"gateway/middleware"
	"gateway/util"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func newReverseProxyMiddleware(configMap map[string]any) (*ReverseProxyMiddleware, error) {
	proxyConfig, err := parseConfig(configMap)
	if err != nil {
		return nil, err
	}

	return &ReverseProxyMiddleware{
		proxyConfig: proxyConfig,
	}, err
}

type ReverseProxyMiddleware struct {
	proxyConfig *ReverseProxyConfiguration
}

func (r *ReverseProxyMiddleware) Handle(ctx *middleware.Context, writer http.ResponseWriter, request *http.Request) error {
	targetURL, err := r.matchPredicates(request)
	if targetURL == nil {
		// 手动设置 HTTP 响应码为 404
		writer.WriteHeader(http.StatusNotFound)
		ctx.Next(writer, request)
		return nil
	}
	proxy, err := r.buildProxy(targetURL)
	if err != nil {
		return err
	}
	proxy.ServeHTTP(writer, request)
	return nil
}

func (r *ReverseProxyMiddleware) buildProxy(targetURL *url.URL) (*httputil.ReverseProxy, error) {
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ErrorLog = util.NewHttpLogger()
	return proxy, nil
}

func (r *ReverseProxyMiddleware) matchPredicates(request *http.Request) (*url.URL, error) {

	for _, proxyInfo := range r.proxyConfig.Proxy {

		mathResult := true

		for _, predicate := range proxyInfo.Predicates {
			parts := strings.Split(predicate, "=")
			if len(parts) != 2 {
				panic(fmt.Sprintf("invalid key-value pair: %s", predicate))
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "After":
			case "Before":
				// 获取请求时间
				requestTime := request.Header.Get("Date")
				// 解析请求时间字符串
				parsedTime, _ := time.Parse(time.RFC1123, requestTime)

				switch key {
				case "After":
					if !timeMatch.After(parsedTime, value) {
						mathResult = false
						break
					}
				case "Before":
					if !timeMatch.Before(parsedTime, value) {
						mathResult = false
						break
					}
				}
			case "Host":
				if !host.Match(request.Method, value) {
					mathResult = false
					break
				}
			case "Method":
				if !method.Match(request.Method, value) {
					mathResult = false
					break
				}
			case "Path":
				if !path.Match(request.URL.Path, value) {
					mathResult = false
					break
				}

			}
		}
		if mathResult {
			targetURL, err := url.Parse(proxyInfo.Uri)
			return targetURL, err
		}
	}
	return nil, nil
}

type ProxyInfo struct {
	Id         string   `config:"id"`
	Uri        string   `config:"uri"`
	Predicates []string `config:"predicates"`
}
type ReverseProxyConfiguration struct {
	Proxy []ProxyInfo `config:"proxy"`
}

func parseConfig(configMap map[string]any) (*ReverseProxyConfiguration, error) {
	//解析配置
	config := &ReverseProxyConfiguration{
		Proxy: make([]ProxyInfo, 0),
	}
	err := util.UnpackConfig(configMap, config)
	return config, err
}

func init() {
	middleware.RegisteredMiddlewares.RegisterHandler("proxy", func(configMap map[string]any) (middleware.Handler, error) {
		return newReverseProxyMiddleware(configMap)
	})
}
