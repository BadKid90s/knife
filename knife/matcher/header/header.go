// 关于请求头的匹配器

package header

import (
	"knife"
	"net/http"
)

func RespExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(response.Header(), key, "")
	}
}

func RespNotExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(response.Header(), key, "")
	}
}

func RespNe(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(response.Header(), key, value)
	}
}

func RespEq(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(response.Header(), key, value)
	}
}

func ReqExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(request.Header, key, "")
	}
}

func ReqNotExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(request.Header, key, "")
	}
}

func ReqNe(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(request.Header, key, value)
	}
}

func ReqEq(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(request.Header, key, value)
	}
}

func equalHeaderValue(header http.Header, key, value string) bool {
	return header.Get(key) == value
}
