package processor

import (
	"fmt"
	"gateway/matcher/host"
	"gateway/matcher/method"
	"gateway/matcher/path"
	"gateway/middleware"
	"gateway/util"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
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

func (r *ReverseProxyMiddleware) Handle(_ *middleware.Context, writer http.ResponseWriter, request *http.Request) error {
	proxy, err := r.buildProxy(request)
	if err != nil {
		return err
	}
	proxy.ServeHTTP(writer, request)
	return nil
}

func (r *ReverseProxyMiddleware) buildProxy(request *http.Request) (*httputil.ReverseProxy, error) {
	targetURL, err := r.matchPredicates(request)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ErrorLog = util.NewHttpLogger()
	return proxy, nil
}

func (r *ReverseProxyMiddleware) matchPredicates(request *http.Request) (*url.URL, error) {

	targetPath := request.URL.Path

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
			case "Path":
				if !path.Match(request.URL.Path, value) {
					mathResult = false
					break
				}

			}
		}
		if mathResult {
			targetPath = proxyInfo.Uri
			break
		}
	}

	targetURL, err := url.Parse(targetPath)
	return targetURL, err
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
