package gc

import (
	"context"

	"google.golang.org/grpc"
)

type StreamContext struct {
	ctx context.Context
	grpc.ServerStream
}

func NewStreamContext(ctx context.Context, ss grpc.ServerStream) *StreamContext {
	return &StreamContext{
		ctx:          ctx,
		ServerStream: ss,
	}
}

func (x *StreamContext) Context() context.Context {
	return x.ctx
}
