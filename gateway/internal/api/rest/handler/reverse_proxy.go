package handler

import (
	"net/http/httputil"

	"github.com/Axot017/k8s-playground/gateway/internal/config"
	"github.com/labstack/echo/v4"
)

type ReverseProxy struct {
	profileProxy *httputil.ReverseProxy
}

func NewReverseProxy(cfg *config.Config) *ReverseProxy {
	profileProxy := httputil.NewSingleHostReverseProxy(&cfg.ProfileServiceUrl)
	return &ReverseProxy{
		profileProxy: profileProxy,
	}
}

func (h *ReverseProxy) Register(r *echo.Group) {
	r.Any("/api/profile/*", echo.WrapHandler(h.profileProxy))
}
