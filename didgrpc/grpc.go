package didgrpc

import (
	"fmt"
	"log"
	"net"
	"os"

	didgrpc "github.com/dlacreme/did-proto"
	"google.golang.org/grpc"
)

func StartServer() {
	fmt.Println("Hello world!")
	host := fmt.Sprintf("%s:%s", os.Getenv("GRPC_HOSTNAME"), os.Getenv("GRPC_PORT"))
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption = make([]grpc.ServerOption, 0)
	server := grpc.NewServer(opts...)
	didgrpc.RegisterDidServer(server, newServer())
	fmt.Printf("gRPC server listening on [%s]", host)
	server.Serve(lis)
}
