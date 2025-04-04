package gcServer

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewInsecureServer(opts ...grpc.ServerOption) (sv *grpc.Server) {
	opts = append(opts, grpc.Creds(insecure.NewCredentials()))
	return grpc.NewServer(opts...)
}
