package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%s --> [%s] %s\n", request.RemoteAddr, request.Method, request.URL.Path)
		next.ServeHTTP(writer, request)
	})
}
