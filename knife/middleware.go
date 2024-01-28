package knife

var AllowTrueMiddlewareMatcher = MiddlewareMatcher(func(HttpResponseWriter, HttpRequest) bool {
	return true
})

type (
	MiddlewareMatcher func(HttpResponseWriter, HttpRequest) bool

	MiddlewareFunc func(*Context)

	Middleware struct {
		matcher MiddlewareMatcher
		handler MiddlewareFunc
	}
)

func NewMiddleware(matcher MiddlewareMatcher, handler MiddlewareFunc) *Middleware {
	return &Middleware{
		matcher: matcher,
		handler: handler,
	}
}
