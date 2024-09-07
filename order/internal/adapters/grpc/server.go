package grpc

import (
	"fmt"
	"github.com/fmo/microservices-book/order/config"
	"github.com/fmo/microservices-book/order/internal/ports"
	"github.com/fmo/microservices-proto/golang/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()

	a.server = grpcServer
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" { // enables reflection to make grpcurl
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
