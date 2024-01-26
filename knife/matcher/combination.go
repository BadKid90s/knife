// 将匹配器进行组合的匹配器

package matcher

import "knife"

// All 所以匹配器都完成返回true,否则false
func All(matches ...knife.MiddlewareMatcher) knife.MiddlewareMatcher {
	return func(writer knife.HttpResponseWriter, request knife.HttpRequest) bool {
		for _, matcher := range matches {
			if !matcher(writer, request) {
				return false
			}
		}
		return true
	}
}

// Any 任意匹配器完成返回true,否则false
func Any(matches ...knife.MiddlewareMatcher) knife.MiddlewareMatcher {
	return func(writer knife.HttpResponseWriter, request knife.HttpRequest) bool {
		for _, matcher := range matches {
			if matcher(writer, request) {
				return true
			}
		}
		return false
	}
}
