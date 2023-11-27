package rest

import (
	"github.com/Axot017/k8s-playground/profile/internal/api/rest/handler"
	"github.com/Axot017/k8s-playground/profile/internal/config"
	"github.com/labstack/echo/v4"
)

type Router struct {
	*echo.Echo
}

func NewRouter(handlers []handler.Handler, config *config.Config) *Router {
	echo := echo.New()
	echo.HideBanner = true
	echo.HidePort = true

	g := echo.Group("/api/profile")

	for _, h := range handlers {
		h.Register(g)
	}
	return &Router{
		Echo: echo,
	}
}
