package api

import (
	"context"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/api/handlers/authors"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/api/handlers/books"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Init() {
	r := mux.NewRouter()
	CORSOptions()
	r.Use(middleware.LoggerMiddleware)
	r.Use(middleware.AuthenticatorMiddleware)
	books.SetRouters(r)
	authors.SetRouters(r)

	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS()(r),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ShutdownServer(srv, time.Second*10)
}

func ShutdownServer(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func CORSOptions() {
	handlers.AllowedOrigins([]string{"*"})
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
}
