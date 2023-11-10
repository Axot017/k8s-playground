package api

import (
	"github.com/Axot017/k8s-playground/profile/internal/api/handler"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	*chi.Mux
}

func NewRouter(handlers []handler.Handler) *Router {
	mux := chi.NewRouter()
	for _, h := range handlers {
		h.Register(mux)
	}
	return &Router{
		Mux: mux,
	}
}
