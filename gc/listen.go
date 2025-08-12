package gc

import (
	"fmt"
	"net"
	"strings"

	"github.com/d3v-friends/go-tools/fnPanic"
	"google.golang.org/grpc"
)

func Listen(sv *grpc.Server, port string) {
	var lis net.Listener

	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	var err error
	if lis, err = net.Listen("tcp", port); err != nil {
		return
	}

	fmt.Printf("🚀 gRPC server ready at localhost%s\n", port)
	fnPanic.On(sv.Serve(lis))
}
