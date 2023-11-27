package tracing

import (
	"context"
	"log/slog"
)

type tracingID string

const tracingIDKey tracingID = "tracing_id"

func SetTracingID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, tracingIDKey, id)
}

func GetTracingID(ctx context.Context) string {
	requestId, _ := ctx.Value(tracingIDKey).(string)
	if requestId == "" {
		slog.ErrorContext(ctx, "tracing_id is empty")
	}

	return requestId
}
