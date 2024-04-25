// 关于请求头的匹配器

package header

import (
	"knife"
	"net/http"
)

func ResponseExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(response.Header(), key, "")
	}
}

func ResponseNotExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(response.Header(), key, "")
	}
}

func ResponseNe(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(response.Header(), key, value)
	}
}

func ResponseEq(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(response.Header(), key, value)
	}
}

func RequestExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(request.Header, key, "")
	}
}

func RequestNotExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(request.Header, key, "")
	}
}

func HeaderRequestNe(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(request.Header, key, value)
	}
}

func HeaderRequestEq(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(request.Header, key, value)
	}
}

func equalHeaderValue(header http.Header, key, value string) bool {
	return header.Get(key) == value
}
