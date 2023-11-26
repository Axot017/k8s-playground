package service

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/Axot017/k8s-playground/gateway/internal/api/rest"
	"github.com/Axot017/k8s-playground/gateway/internal/config"
)

type Http struct {
	server *http.Server
}

func NewHttp(mux *rest.Router, config *config.Config) *Http {
	slog.Info("Creating HTTP server", "port", config.Port)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: mux,
	}
	return &Http{
		server: &server,
	}
}

func (h *Http) Serve(ctx context.Context) error {
	slog.InfoContext(ctx, "Starting HTTP server")
	listener, err := net.Listen("tcp", h.server.Addr)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to start HTTP server", "error", err)
		return err
	}
	go func() {
		_ = h.server.Serve(listener)
	}()

	return nil
}

func (h *Http) Shutdown(ctx context.Context) error {
	slog.InfoContext(ctx, "Shutting down HTTP server")
	return h.server.Shutdown(ctx)
}
