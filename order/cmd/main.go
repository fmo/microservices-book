package main

import (
	"github.com/fmo/microservices-book/order/config"
	"github.com/fmo/microservices-book/order/internal/adapters/db"
	"github.com/fmo/microservices-book/order/internal/adapters/grpc"
	"github.com/fmo/microservices-book/order/internal/application/core/api"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	service     = "order"
	environment = "dev"
	id          = 1
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
