package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/zerobase-xyz/go-grpc-sample/pb"

	"google.golang.org/grpc"
)

var (
	empty = new(pb.Empty)
)

type HostnameService struct{}

func factorize(number int, c chan<- []int) {
	var a []int
	for i := 1; i < number+1; i++ {
		if number%i == 0 {
			a = append(a, i)
		}
		c <- a
	}
	close(c)
}

func (s *HostnameService) GetPodHostname(ctx context.Context, in *pb.Empty) (*pb.PodHostnameResponse, error) {
	name, _ := os.Hostname()
	//number := rand.Intn(999999) + 1000000
	//c := make(chan []int)
	//go factorize(number, c)

	fmt.Printf("%v\n", name)
	//for i := range c {
	//	fmt.Printf("result: %d\n", i)
	//}

	return &pb.PodHostnameResponse{
		Name: name,
	}, nil
}

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	HostnameService := &HostnameService{}
	pb.RegisterHostnamePodServiceServer(server, HostnameService)
	server.Serve(listenPort)
}
