package template

import (
	"context"
	"github.com/d3v-friends/go-pure/fnLogger"
	"github.com/d3v-friends/go-pure/fnPanic"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"time"
)

type FnInterceptor func(ctx context.Context) (res context.Context, err error)

func Interceptor(fn FnInterceptor, logger fnLogger.IfLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var requestAt = time.Now()
		var requestLogger = logger.WithFields(fnLogger.Fields{
			"requestId": uuid.NewString(),
			"method":    info.FullMethod,
			"requestAt": time.Now(),
		})
		requestLogger.Trace("requested")

		ctx = fnPanic.OnValue(fn(ctx))

		defer func() {
			var responseAt = time.Now()
			var responseLogger = requestLogger.
				WithFields(fnLogger.Fields{
					"responseAt": responseAt,
					"durations":  responseAt.UnixMilli() - requestAt.UnixMilli(),
				})

			if err == nil {
				responseLogger.
					Trace("responded")
			} else {
				responseLogger.
					Error("err=%s", err.Error())
			}
		}()

		return handler(ctx, req)
	}
}
