package knife

import "net/http"

type BeforeFunc func(writer HttpResponseWriter)

type HttpResponseWriter interface {
	http.ResponseWriter
	Status() int
	Written() bool
	Size() int
	Before(beforeFunc BeforeFunc)
}

func NewResponseWriter(writer http.ResponseWriter) HttpResponseWriter {
	return &responseWriter{ResponseWriter: writer}
}

type responseWriter struct {
	http.ResponseWriter
	pendingStatus       int
	status              int
	size                int
	beforeFunctions     []BeforeFunc
	callingBeforeStatus bool
}

func (w *responseWriter) WriteHeader(s int) {
	if w.Written() {
		return
	}
	w.pendingStatus = s
	w.callBefore()
	if w.Written() {
		return
	}
	w.status = s
	w.ResponseWriter.WriteHeader(s)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if !w.Written() {
		w.WriteHeader(http.StatusOK)
	}
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}

func (w *responseWriter) Status() int {
	if w.Written() {
		return w.status
	}

	return w.pendingStatus
}

func (w *responseWriter) Size() int {
	return w.size
}

func (w *responseWriter) Written() bool {
	return w.status != 0
}

func (w *responseWriter) Before(before BeforeFunc) {
	w.beforeFunctions = append(w.beforeFunctions, before)
}

func (w *responseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

func (w *responseWriter) callBefore() {
	if w.callingBeforeStatus {
		return
	}

	w.callingBeforeStatus = true
	defer func() { w.callingBeforeStatus = false }()

	for i := len(w.beforeFunctions) - 1; i >= 0; i-- {
		w.beforeFunctions[i](w)
	}
}
