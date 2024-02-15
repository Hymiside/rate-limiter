package handler

import (
	"net/http"
	"time"
)

func (h *Handler) rateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if h.limiter.Requests[ip] >= h.limiter.MaxRequests {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		h.limiter.Requests[ip]++
		go func() {
			time.Sleep(h.limiter.Interval)
			h.limiter.Requests[ip]--
		}()
		next.ServeHTTP(w, r)
	})
}