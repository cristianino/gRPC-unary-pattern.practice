package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"example.grpc.com/hellopb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	fmt.Printf("Hello function was called with %v \n", req)
	firstName := req.GetHello().GetFirstName()
	prefix := req.GetHello().GetPrefix()

	customHello := fmt.Sprintf("Welcome %v %v!", prefix, firstName)
	res := &hellopb.HelloResponse{
		CustomHello: customHello,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello, Go server is runnig")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Fail to listen %v", err)
	}

	s := grpc.NewServer()

	hellopb.RegisterHelloServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
