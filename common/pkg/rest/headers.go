package rest

import (
	"net/http"

	"github.com/Axot017/k8s-playground/common/pkg/tracing"
)

const RequestIDHeader = "X-Request-ID"

func RequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get(RequestIDHeader)

		if requestId == "" {
			tracing.SetRequestId(r.Context(), requestId)
		}

		h.ServeHTTP(w, r)
	})
}
