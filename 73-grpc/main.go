package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"grpc-demo/grpc-hello-world/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// server is used to implement helloworld.GreeterServer
type server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func (s *server) GetRandom(ctx context.Context, int *emptypb.Empty) (out *helloworld.RandomReply, err error) {
	//log.Printf("Received: %v", in.GetName())
	return &helloworld.RandomReply{Rand: int64(rand.Intn(100))}, nil
}

func (s *server) GetRandomStream(in *emptypb.Empty, out grpc.ServerStreamingServer[helloworld.RandomReply]) error {
	for {
		time.Sleep(time.Second * 2)
		err := out.Send(&helloworld.RandomReply{Rand: int64(rand.Intn(100))})
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	log.Println("gRPC server is listening on port 50052")

	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
