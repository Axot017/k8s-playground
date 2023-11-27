package rest

import (
	"log/slog"

	"github.com/Axot017/k8s-playground/common/pkg/tracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Recover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(
		middleware.RecoverConfig{
			LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
				slog.ErrorContext(c.Request().Context(), "Unexpected critical error", "error", err)

				return err
			},
		},
	)
}

func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestIDHeader := c.Request().Header.Get(RequestIDHeader)
			if requestIDHeader == "" {
				return next(c)
			}

			ctx := tracing.SetTracingID(c.Request().Context(), requestIDHeader)
			c.Request().WithContext(ctx)

			return next(c)
		}
	}
}

func Logger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogError:   true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			if values.Error == nil {
				slog.DebugContext(
					c.Request().Context(),
					"Request",
					"method",
					values.Method,
					"uri",
					values.URI,
					"status",
					values.Status,
					"duration",
					values.Latency,
				)
			} else {
				if values.Status < 500 {
					slog.WarnContext(
						c.Request().Context(),
						"Request",
						"method",
						values.Method,
						"uri",
						values.URI,
						"status",
						values.Status,
						"error",
						values.Error,
						"duration",
						values.Latency,
					)
				} else {
					slog.ErrorContext(
						c.Request().Context(),
						"Request",
						"method",
						values.Method,
						"uri",
						values.URI,
						"status",
						values.Status,
						"error",
						values.Error,
						"duration",
						values.Latency,
					)
				}
			}

			return nil
		},
	})
}
