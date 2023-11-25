package api

import (
	"github.com/Axot017/k8s-playground/gateway/internal/api/handler"
	"github.com/Axot017/k8s-playground/gateway/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	chi.Router
}

func NewRouter(handlers []handler.Handler, logger *config.Config) *Router {
	router := chi.NewMux()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)

	if logger.Debug {
		router.Mount("/debug", middleware.Profiler())
	}

	for _, h := range handlers {
		h.Register(router)
	}
	return &Router{
		Router: router,
	}
}
