package matcher

import "strings"

func Path(path, formatter string) bool {
	return strings.HasPrefix(path, formatter)
}
