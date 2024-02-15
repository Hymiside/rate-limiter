package handler

import (
	"net/http"
	"time"
)

type RateLimiter struct {
	Interval time.Duration
	MaxRequests int
	Requests map[string]int
}

type Handler struct {
	limiter *RateLimiter
}

func NewHandler() *Handler {
	return &Handler{limiter: &RateLimiter{
		Interval: 5 * time.Second,
		MaxRequests: 10,
		Requests: make(map[string]int),
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