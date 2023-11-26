package context

import (
	"context"
	"log/slog"
)

type RequestId string

const requestIdKey RequestId = "request_id"

func SetRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestIdKey, requestId)
}

func GetRequestId(ctx context.Context) string {
	requestId, _ := ctx.Value(requestIdKey).(string)
	if requestId == "" {
		slog.Error("requestId is empty")
	}

	return requestId
}
