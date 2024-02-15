package handler

import (
	"net/http"
	"time"
)

type RateLimiter struct {
	interval time.Duration
	maxRequests int
	requests map[string]int
}

type Handler struct {
	limiter *RateLimiter
}

func NewHandler() *Handler {
	return &Handler{limiter: &RateLimiter{
		interval: 5 * time.Second,
		maxRequests: 10,
		requests: make(map[string]int),
	}}
}

func (h *Handler) InitHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", h.rateLimiterMiddleware(http.HandlerFunc(h.helloWorld)))
	return mux
}

func (h *Handler) helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}