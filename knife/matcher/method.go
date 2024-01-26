// 关于请求方法的匹配器

package matcher

import (
	"knife"
	"strings"
)

func MethodEq(method string) knife.MiddlewareMatcher {
	method = strings.ToUpper(method)
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return request.Method == method
	}
}
func MethodNe(method string) knife.MiddlewareMatcher {
	method = strings.ToUpper(method)
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return request.Method != method
	}
}

func MethodAny(methods ...string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		for _, method := range methods {
			method = strings.ToUpper(method)
			if request.Method == method {
				return true
			}
		}
		return false
	}
}
