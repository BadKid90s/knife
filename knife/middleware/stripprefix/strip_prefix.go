package stripprefix

import (
	"knife"
	"net/url"
	"strings"
)

// StripPrefix 剥离前缀
func StripPrefix(prefix string) knife.MiddlewareFunc {
	return func(context *knife.Context) {
		r := context.Req
		p := strings.TrimPrefix(r.URL.Path, prefix)
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r.URL = new(url.URL)
			r.URL.Path = p
			r.URL.RawPath = rp
		}
		context.Next()
	}
}
