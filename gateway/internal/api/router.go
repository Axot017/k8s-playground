package api

import (
	"net/http"

	"github.com/Axot017/k8s-playground/gateway/internal/api/handler"
)

type Router struct {
	*http.ServeMux
}

func NewRouter(handlers []handler.Handler) *Router {
	mux := http.NewServeMux()
	for _, h := range handlers {
		h.Register(mux)
	}
	return &Router{
		ServeMux: mux,
	}
}
