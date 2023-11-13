package internal

import (
	"context"

	"go.uber.org/fx"

	"github.com/Axot017/k8s-playground/gateway/internal/api"
	"github.com/Axot017/k8s-playground/gateway/internal/config"
	"github.com/Axot017/k8s-playground/gateway/internal/service"
)

func StartApp() {
	fx.New(
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
