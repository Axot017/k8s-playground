package api

import (
	"github.com/Axot017/k8s-playground/profile/internal/api/handler"
	"go.uber.org/fx"
)

func Providers() []interface{} {
	return []interface{}{
		fx.Annotate(
			NewRouter,
			fx.ParamTags(`group:"handlers"`),
		),
		fx.Annotate(
			handler.NewHealth,
			fx.ResultTags(`group:"handlers"`),
			fx.As(new(handler.Handler)),
		),
	}
}