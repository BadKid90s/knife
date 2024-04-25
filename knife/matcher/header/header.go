// 关于请求头的匹配器

package header

import (
	"knife"
	"net/http"
)

func HeaderResponseExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(response.Header(), key, "")
	}
}

func HeaderResponseNotExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(response.Header(), key, "")
	}
}

func HeaderResponseNe(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(response.Header(), key, value)
	}
}

func HeaderResponseEq(key, value string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return equalHeaderValue(response.Header(), key, value)
	}
}

func HeaderRequestExists(key string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return !equalHeaderValue(request.Header, key, "")
	}
}

func HeaderRequestNotExists(key string) knife.MiddlewareMatcher {
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
