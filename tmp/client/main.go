package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/zerobase-xyz/go-grpc-sample/pb"

	"google.golang.org/grpc"
)

var (
	empty = new(pb.Empty)
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewHostnamePodServiceClient(conn)
	res, err := client.GetPodHostname(context.TODO(), empty)
	fmt.Printf("result:%#v \n", res.Name)
	fmt.Printf("error::%#v \n", err)
}
