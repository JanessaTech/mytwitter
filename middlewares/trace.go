package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hi-supergirl/mytwitter/logging"
)

var RequestIdKeyInHeader = "X-Request-ID"
var requestIdKey = "requestId"

func RequestTraceMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get(RequestIdKeyInHeader)
		if requestId == "" {
			requestId = uuid.New().String()
		}
		ctx.Request = ctx.Request.WithContext(contextWithRequestId(ctx, requestId))
		logger := logging.DefaultLogger().With(requestIdKey, requestId)
		ctx.Request = ctx.Request.WithContext(logging.WithLogger(ctx, logger))
		ctx.Writer.Header().Set(RequestIdKeyInHeader, requestId)
	}
}

func contextWithRequestId(ctx context.Context, requestId string) context.Context {
	if gCtx, ok := ctx.(*gin.Context); ok {
		ctx = gCtx.Request.Context()
	}
	return context.WithValue(ctx, requestIdKey, requestId)
}
