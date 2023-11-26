package middleware

import (
	"net/http"

	"github.com/Axot017/k8s-playground/common/pkg/context"
	"github.com/Axot017/k8s-playground/common/pkg/rest"
	"github.com/segmentio/ksuid"
)

func RequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := ksuid.New().String()

		ctx := context.SetRequestId(r.Context(), requestId)
		r = r.WithContext(ctx)

		r.Header.Set(rest.RequestIDHeader, requestId)

		h.ServeHTTP(w, r)
	})
}
