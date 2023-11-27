package rest

import (
	"github.com/Axot017/k8s-playground/common/pkg/rest"
	"github.com/Axot017/k8s-playground/gateway/internal/api/rest/handler"
	"github.com/Axot017/k8s-playground/gateway/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/segmentio/ksuid"
)

type Router struct {
	*echo.Echo
}

func NewRouter(handlers []handler.Handler, config *config.Config) *Router {
	echo := echo.New()
	echo.HideBanner = true
	echo.HidePort = true
	echo.Debug = config.Debug

	echo.Use(rest.Recover())
	echo.Use(rest.Logger())
	echo.Use(middleware.TimeoutWithConfig(
		middleware.TimeoutConfig{},
	))
	echo.Use(middleware.RequestIDWithConfig(
		middleware.RequestIDConfig{
			Generator: func() string {
				return ksuid.New().String()
			},
			TargetHeader: rest.RequestIDHeader,
		},
	))

	g := echo.Group("")

	for _, h := range handlers {
		h.Register(g)
	}
	return &Router{
		Echo: echo,
	}
}
