package main

import (
	"context"
	"flag"
	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pristine/handlers"
	"pristine/providers"
	"pristine/repositories"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	idGen := providers.NewIdGenerator()
	authProvider := providers.NewAuthProvider()
	orgDbRepo := repositories.NewOrgDbRepo(idGen)

	authHandler := handlers.NewAuthHandler(orgDbRepo, authProvider)

	r := mux.NewRouter()
	r.Use(handlers.LoggingMiddleware)
	r.HandleFunc("/api/health", handlers.HealthCheckHandler)
	r.HandleFunc("/api/login", authHandler.LoginHandler)
	r.PathPrefix("/").Handler(handlers.NewSpaHandler("build", "index.html"))

	srv := &http.Server{
		Handler:      muxHandlers.RecoveryHandler()(r),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("starting server: %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
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
