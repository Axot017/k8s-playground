package internal

import (
	"context"

	"github.com/Axot017/k8s-playground/profile/internal/api"
	"github.com/Axot017/k8s-playground/profile/internal/config"
	"github.com/Axot017/k8s-playground/profile/internal/service"
	"go.uber.org/fx"
)

func StartApp() {
	fx.New(
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
