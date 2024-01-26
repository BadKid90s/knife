package knife

var AllowTrueMiddlewareMatcher = MiddlewareMatcher(func(HttpResponseWriter, HttpRequest) bool {
	return true
})

type (
	MiddlewareMatcher func(HttpResponseWriter, HttpRequest) bool

	MiddlewareFunc func(*Context)
)
