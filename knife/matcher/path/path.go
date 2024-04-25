// 关于路径的匹配器

package path

import (
	"fmt"
	"knife"
	"log"
	"regexp"
	"strings"
)

// Prefix 路径前缀匹配
func Prefix(prefix string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		return strings.HasPrefix(request.URL.Path, prefix)
	}
}

// Match 路径正则匹配
func Match(pattern string) knife.MiddlewareMatcher {
	return func(response knife.HttpResponseWriter, request knife.HttpRequest) bool {
		matched, err := isPathMatched(request.URL.Path, pattern)
		if err != nil {
			log.Println("path matched failed")
			return false
		}
		return matched
	}
}

// isPathMatched 路径匹配方式
// path 路径
// pattern 正则表达式
func isPathMatched(path string, pattern string) (bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, fmt.Errorf("invalid regular expression: %w", err)
	}
	return re.MatchString(path), nil
}
