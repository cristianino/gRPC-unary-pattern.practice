package main

import (
	"context"
	"fmt"
	"log"

	"example.grpc.com/hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Go client is running")

	cc, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Fail to connect %v", err)
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)
	helloUnary(c)
}

func helloUnary(c hellopb.HelloServiceClient) {
	fmt.Println("starting unary RPC Hello")

	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "Cristian",
			Prefix:    "Sr",
		},
	}

	res, err := c.Hello(context.Background(), req)

	if err != nil {
		log.Fatalf("error, calling Hello RPC: %v", err)
	}

	log.Printf("Response Hello: %v", res.CustomHello)
}
