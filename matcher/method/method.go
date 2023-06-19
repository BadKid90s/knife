package method

import "strings"

func Match(method, patterns string) bool {
	if len(patterns) == 0 || patterns == "*" {
		return true
	}
	for _, s := range strings.Split(patterns, ",") {
		if strings.ToLower(method) == strings.ToLower(s) {
			return true
		}
	}
	return strings.ToLower(method) == strings.ToLower(patterns)
}
