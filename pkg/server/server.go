package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(ctx context.Context, handler http.Handler) error {
	httpServer := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		httpServer.Shutdown(ctx)
	}(ctx)

	log.Printf("authentication microservice launched on http://%s:%s/", "localhost", "8080")
	return httpServer.ListenAndServe()
}