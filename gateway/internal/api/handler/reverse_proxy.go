package handler

import (
	"net/http/httputil"

	"github.com/Axot017/k8s-playground/gateway/internal/config"
	"github.com/go-chi/chi/v5"
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

func (h *ReverseProxy) Register(r chi.Router) {
	r.Mount("/api/profile", h.profileProxy)
}
