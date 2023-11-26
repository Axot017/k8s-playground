package tracing

import (
	"context"
	"log/slog"
)

type requestId string

const requestIdKey requestId = "request_id"

func SetRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestIdKey, requestId)
}

func GetRequestId(ctx context.Context) string {
	requestId, _ := ctx.Value(requestIdKey).(string)
	if requestId == "" {
		slog.ErrorContext(ctx, "requestId is empty")
	}

	return requestId
}
