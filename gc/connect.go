package gc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(host string) (*grpc.ClientConn, error) {
	return grpc.NewClient(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}
