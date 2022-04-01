package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/kevinmichaelchen/some-protos-go/pkg/some/protos/greeter/v1beta1"
	"google.golang.org/grpc"
)

type Server struct{}

func (s Server) SayHello(ctx context.Context, req *v1beta1.SayHelloRequest) (*v1beta1.SayHelloResponse, error) {
	log.Println("SayHello was invoked")
	return &v1beta1.SayHelloResponse{Message: "foobar"}, nil
}

func main() {
	server := grpc.NewServer()
	srv := Server{}
	v1beta1.RegisterGreeterServiceServer(server, srv)
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Starting server...")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
