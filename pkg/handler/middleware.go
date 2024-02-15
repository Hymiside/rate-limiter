package handler

import (
	"net/http"
	"time"
)

func (h *Handler) rateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if h.limiter.requests[ip] >= h.limiter.maxRequests {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		h.limiter.requests[ip]++
		go func() {
			time.Sleep(h.limiter.interval)
			h.limiter.requests[ip]--
		}()
		next.ServeHTTP(w, r)
	})
}