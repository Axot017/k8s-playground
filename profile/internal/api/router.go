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
	mux.Route("/api/profile", func(r chi.Router) {
		for _, h := range handlers {
			h.Register(r)
		}
	})
	return &Router{
		Mux: mux,
	}
}
