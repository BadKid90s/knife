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
