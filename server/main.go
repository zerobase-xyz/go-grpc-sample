package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/zerobase-xyz/go-grpc-sample/pb"

	"google.golang.org/grpc"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var (
	empty = new(pb.Empty)
)

type HostnameService struct{}

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
	addr := net.JoinHostPort(
		os.Getenv("DD_AGENT_HOST"),
		os.Getenv("DD_TRACE_AGENT_PORT"),
	)
	tracer.Start(tracer.WithAnalytics(true), tracer.WithAgentAddr(addr), tracer.WithGlobalTag("env", "dev"))
	defer tracer.Stop()
	si := grpctrace.StreamServerInterceptor()
	ui := grpctrace.UnaryServerInterceptor()

	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer(grpc.StreamInterceptor(si), grpc.UnaryInterceptor(ui))
	HostnameService := &HostnameService{}
	pb.RegisterHostnamePodServiceServer(server, HostnameService)
	server.Serve(listenPort)
}
