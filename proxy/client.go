package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/zerobase-xyz/go-grpc-sample/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	empty := new(pb.Empty)
	conn, err := grpc.Dial("grpc-server.poc.dmmga.me:443", grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	for {
		client := pb.NewHostnamePodServiceClient(conn)
		res, err := client.GetPodHostname(context.TODO(), empty)
		if err != nil {
			log.Fatal("error:", err)
			break
		}
		fmt.Printf("result:%#v \n", res.Name)
	}
}
