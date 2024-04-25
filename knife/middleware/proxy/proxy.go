package proxy

import (
	"knife"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Proxy(target string) knife.MiddlewareFunc {
	return func(c *knife.Context) {
		serveReverseProxy(target, c.Writer, c.Req)
		c.Next()
	}
}

// ServeReverseProxy Serve a reverse proxy for a given url
func serveReverseProxy(target string, writer http.ResponseWriter, req *http.Request) {
	// parse the url
	remote, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(request *http.Request) {
		targetQuery := remote.RawQuery
		request.URL.Scheme = remote.Scheme
		request.URL.Host = remote.Host
		request.Host = remote.Host // todo 这个是关键
		request.URL.Path, request.URL.RawPath = joinURLPath(remote, request.URL)

		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}

		if targetQuery == "" || request.URL.RawQuery == "" {
			request.URL.RawQuery = targetQuery + request.URL.RawQuery
		} else {
			request.URL.RawQuery = targetQuery + "&" + request.URL.RawQuery
		}
	}

	// 修改响应头
	proxy.ModifyResponse = func(response *http.Response) error {
		response.Header.Add("Access-Control-Allow-Origin", "*")
		return nil
	}

	// Note that ServeHttp is not blocking and uses a go routine under the hood
	proxy.ServeHTTP(writer, req)
}

// go sdk 源码
func singleJoiningSlash(a, b string) string {
	aSlash := strings.HasSuffix(a, "/")
	bSlash := strings.HasPrefix(b, "/")
	switch {
	case aSlash && bSlash:
		return a + b[1:]
	case !aSlash && !bSlash:
		return a + "/" + b
	}
	return a + b
}

// go sdk 源码
func joinURLPath(a, b *url.URL) (path, rawPath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	aPath := a.EscapedPath()
	bPath := b.EscapedPath()

	aSlash := strings.HasSuffix(aPath, "/")
	bSlash := strings.HasPrefix(bPath, "/")

	switch {
	case aSlash && bSlash:
		return a.Path + b.Path[1:], aPath + bPath[1:]
	case !aSlash && !bSlash:
		return a.Path + "/" + b.Path, aPath + "/" + bPath
	}
	return a.Path + b.Path, aPath + bPath
}
