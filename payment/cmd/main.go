package main

import (
	"github.com/fmo/microservices-book/payment/config"
	"github.com/fmo/microservices-book/payment/internal/adapters/db"
	"github.com/fmo/microservices-book/payment/internal/adapters/grpc"
	"github.com/fmo/microservices-book/payment/internal/application/core/api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
