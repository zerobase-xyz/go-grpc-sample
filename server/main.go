package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/zerobase-xyz/go-grpc-sample/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

var (
	empty = new(pb.Empty)
)

type HostnameService struct{}

type healthServer struct{}

func (h *healthServer) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (h *healthServer) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "service watch is not implemented current version.")
}

func factorize(number int) {
	var a []int
	for i := 1; i < number+1; i++ {
		if number%i == 0 {
			a = append(a, i)
		}
	}
}

func (s *HostnameService) GetPodHostname(ctx context.Context, in *pb.Empty) (*pb.PodHostnameResponse, error) {
	name, _ := os.Hostname()
	number := 10
	factorize(number)

	fmt.Printf("%v\n", name)
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
	health.RegisterHealthServer(server, &healthServer{})
	server.Serve(listenPort)
}
