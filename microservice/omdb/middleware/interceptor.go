package middleware

import (
	"context"
	"omdb/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func InterceptorWithDBLogger(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		nextCtx := ctx
		if ok {
			if values, ok := md[RequestIDKey]; ok && len(values) == 1 {
				nextCtx = context.WithValue(ctx, RequestIDKey, values[0])
			}
		}
		go log.Log("[method]:", info.FullMethod, "[metadata]:", md, "[request]", req)
		return handler(nextCtx, req)
	}
}
