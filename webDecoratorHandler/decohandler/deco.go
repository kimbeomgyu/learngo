package decohandler

import "net/http"

// DecoratorFunc is DecoratorFunc
type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

// DecoHandler is struct
type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler
}

func (deco *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	deco.fn(w, r, deco.h)
}

// NewDecoHandler is NewDecoHandler
func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h:  h,
	}
}
