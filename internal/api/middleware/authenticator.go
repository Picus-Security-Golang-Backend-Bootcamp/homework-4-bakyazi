package middleware

import (
	"net/http"
	"strings"
)

func AuthenticatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			next.ServeHTTP(writer, request)
			return
		}
		authorization := request.Header.Get("Authorization")
		if strings.HasPrefix(authorization, "Bearer ") {
			token := strings.TrimPrefix(authorization, "Bearer ")
			if token != "" {
				next.ServeHTTP(writer, request)
				return
			}
		}
		http.Error(writer, "Token not found", http.StatusUnauthorized)
	})
}
