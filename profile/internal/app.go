package internal

import (
	"context"
	"log/slog"
	"os"

	"github.com/Axot017/k8s-playground/common/pkg/fxlogger"
	"github.com/Axot017/k8s-playground/profile/internal/api"
	"github.com/Axot017/k8s-playground/profile/internal/config"
	"github.com/Axot017/k8s-playground/profile/internal/service"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func StartApp() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	fx.New(
		fx.WithLogger(func() fxevent.Logger { return fxlogger.New(slog.Default()) }),
		fx.Provide(config.Providers()...),
		fx.Provide(service.Providers()...),
		fx.Provide(api.Providers()...),
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
