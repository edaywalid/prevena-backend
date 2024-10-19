package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/config"
)

func isAllowedOrigin(origin string, allowedOrigins []string) bool {
	for _, allowed := range allowedOrigins {
		if origin == allowed {
			return true
		}
	}
	return false
}

func CorsMiddleware(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env := os.Getenv("ENV")

		var allowedOrigins []string
		if env == "production" {
			allowedOrigins = strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
		} else {
			allowedOrigins = []string{"*"}
		}

		origin := r.Header.Get("Origin")

		if env != "production" || isAllowedOrigin(origin, allowedOrigins) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
