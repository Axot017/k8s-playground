package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Axot017/k8s-playground/profile/internal/api/rest"
	"github.com/Axot017/k8s-playground/profile/internal/config"
)

type Http struct {
	router *rest.Router
	config *config.Config
}

func NewHttp(router *rest.Router, config *config.Config) *Http {
	slog.Info("Creating HTTP server", "port", config.Port)
	return &Http{
		router: router,
		config: config,
	}
}

func (h *Http) Serve(ctx context.Context) error {
	slog.InfoContext(ctx, "Starting HTTP server")
	address := fmt.Sprintf(":%d", h.config.Port)
	go func() {
		err := h.router.Start(address)
		if err != nil {
			slog.ErrorContext(ctx, "Error starting HTTP server", "error", err)
		}
	}()

	return nil
}

func (h *Http) Shutdown(ctx context.Context) error {
	slog.InfoContext(ctx, "Shutting down HTTP server")
	return h.router.Shutdown(ctx)
}
