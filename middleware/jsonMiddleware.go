package middleware

import (
	"net/http"
	"strings"
)

func Adapt(handler *http.Handler) http.Handler {
	return AdaptFunctional((*handler).ServeHTTP)
}

func AdaptFunctional(handle http.HandlerFunc) http.HandlerFunc {
	hf := func(rw http.ResponseWriter, request *http.Request) {
		if !strings.HasPrefix(request.Header.Get("Content-Type"), "application/json") {
			http.Error(rw, "Content type not supported (expected application/json)", http.StatusUnsupportedMediaType)
			return
		}

		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		handle(rw, request)
	}

	return http.HandlerFunc(hf)
}
