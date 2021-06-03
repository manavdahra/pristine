package main

import (
	"context"
	"flag"
	"fmt"
	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"pristine/config"
	"pristine/handlers"
	"pristine/providers"
	"pristine/repositories"
	"pristine/services"
	"time"
)

func main() {
	logger := providers.InitLogger()

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	cfg := config.Init()
	idGen := providers.NewIdGenerator()
	mongo := providers.NewMongoDAL(cfg, logger)
	userDbRepo := repositories.NewUserDbRepo(idGen, mongo, logger)
	userService := services.NewUserService(userDbRepo)
	//orgDbRepo := repositories.NewOrgDbRepo(idGen, mongo, logger)

	middlewares := handlers.NewMiddlewares(logger)
	authHandler := handlers.NewAuthHandler(userService, cfg, logger)

	r := mux.NewRouter()

	r.Use(middlewares.LoggingMiddleware)
	r.HandleFunc("/api/health", handlers.HealthCheckHandler)
	r.HandleFunc("/api/signIn", authHandler.LoginHandler)
	r.HandleFunc("/api/signOut", authHandler.LogoutHandler)
	r.PathPrefix("/").Handler(handlers.NewSpaHandler("build", "index.html"))

	//handler := muxHandlers.RecoveryHandler()(r)
	handler := muxHandlers.CORS(
		muxHandlers.AllowedOrigins(cfg.AllowedOrigins()),
		muxHandlers.AllowedHeaders(cfg.AllowedHeaders()),
		muxHandlers.AllowedMethods(cfg.AllowedMethods()),
	)(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("%s:%s", cfg.GetHost(), cfg.GetPort()),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logger.Infow("starting server", "address", srv.Addr, "handler", srv.Handler)
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err)
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
	logger.Info("shutting down")
	logger.Sync()

	os.Exit(0)
}
