package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/dhucsik/onelab-hw4/config"
	_ "github.com/dhucsik/onelab-hw4/docs"
	"github.com/dhucsik/onelab-hw4/service"
	"github.com/dhucsik/onelab-hw4/storage"
	"github.com/dhucsik/onelab-hw4/transport/http"
	"github.com/dhucsik/onelab-hw4/transport/http/handler"
)

// @title Transactions microservice
// @version 1.0
// @description Microservice for storing and managing transactions
// @termOfService http://swagger.io/terms/

// @contact.name Dias Galikhanov
// @contact.email diasgalikhanov@gmail.com

// @host localhost:8587
// @BasePath /api/v1

// Transaction CRUD + Swagger

func main() {
	log.Fatalln(run())
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		return err
	}

	stg, err := storage.New(ctx, conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	svc, svcErr := service.NewManager(stg)
	if svcErr != nil {
		return svcErr
	}

	h := handler.NewManager(svc)
	HTTPServer := http.NewServer(conf, h)

	return HTTPServer.StartHTTPServer(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
