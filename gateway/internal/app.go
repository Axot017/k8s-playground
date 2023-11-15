package internal

import (
	"context"
	"log/slog"
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/Axot017/k8s-playground/common/pkg/fxlogger"
	"github.com/Axot017/k8s-playground/gateway/internal/api"
	"github.com/Axot017/k8s-playground/gateway/internal/config"
	"github.com/Axot017/k8s-playground/gateway/internal/service"
)

func StartApp() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	fx.New(
		fx.WithLogger(func() fxevent.Logger { return fxlogger.New(slog.Default()) }),
		fx.Provide(api.Providers()...),
		fx.Provide(service.Providers()...),
		fx.Provide(config.Providers()...),
		fx.Invoke(startHttpListener),
	).Run()
}

func startHttpListener(lifecycle fx.Lifecycle, httpService *service.Http) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return httpService.Serve(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return httpService.Shutdown(ctx)
		},
	})
}
