package rest

import (
	commonRest "github.com/Axot017/k8s-playground/common/pkg/rest"
	"github.com/Axot017/k8s-playground/profile/internal/api/rest/handler"
	"github.com/Axot017/k8s-playground/profile/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	chi.Router
}

func NewRouter(handlers []handler.Handler, config *config.Config) *Router {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(commonRest.RequestID)

	if config.Debug {
		router.Mount("/debug", middleware.Profiler())
	}

	router.Route("/api/profile", func(r chi.Router) {
		for _, h := range handlers {
			h.Register(r)
		}
	})

	return &Router{
		Router: router,
	}
}
