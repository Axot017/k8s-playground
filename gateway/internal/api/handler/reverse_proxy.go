package handler

import (
	"net/http"
	"net/http/httputil"

	"github.com/Axot017/k8s-playground/gateway/internal/config"
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

func (h *ReverseProxy) Register(r *http.ServeMux) {
	r.Handle("/api/profile/", h.profileProxy)
}
