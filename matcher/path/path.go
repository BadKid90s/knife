package path

import (
	"log"
	osPath "path"
	"strings"
)

func Match(path, patterns string) bool {
	if len(patterns) == 0 || patterns == "/" {
		return true
	}
	for _, pattern := range strings.Split(patterns, ",") {
		if mathPath(path, pattern) {
			return true
		}
	}
	return false
}
func mathPath(path, pattern string) bool {
	match, err := osPath.Match(pattern, path)
	if err != nil {
		log.Printf("match path error path：%v, pattern: %v", path, pattern)
	}
	return match
}
