package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hymiside/rate-limiter/pkg/handler"
	"github.com/Hymiside/rate-limiter/pkg/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handlers := handler.NewHandler()
	mux := handlers.InitHandler()

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		select {
		case <-quit:
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	server := server.NewServer()
	if err := server.StartServer(ctx, mux); err != nil {
		log.Fatal(err)
	}
}
