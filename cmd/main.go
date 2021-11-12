package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/shellingford330/auth/injector/grpc"
	"github.com/shellingford330/auth/injector/http"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// DI
	httpServer, err := http.InitServer()
	if err != nil {
		panic(err)
	}
	grpcServer, err := grpc.InitServer()
	if err != nil {
		panic(err)
	}

	// Start server...
	go func() {
		if err := httpServer.Start(); err != nil {
			panic(err)
		}
	}()
	go func() {
		log.Println("gPRC server running ...")
		if err := grpcServer.Start(); err != nil {
			panic(err)
		}
	}()

	// User intercept
	<-ctx.Done()
	if err := httpServer.GracefulShutdown(); err != nil {
		panic(err)
	}
	log.Println("graceful shutdown http server")
	if err := grpcServer.GracefulShutdown(); err != nil {
		panic(err)
	}
	log.Println("graceful shutdown grpc server")
}
