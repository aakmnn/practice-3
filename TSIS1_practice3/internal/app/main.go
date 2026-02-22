package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"golang/internal/handler"
	"golang/internal/repository"
	"golang/internal/repository/_postgres"
	"golang/internal/usecase"
	"golang/pkg/modules"
)

const apiKey = "my-secret-key"

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := initPostgreConfig()
	pg := _postgres.NewPGXDialect(ctx, cfg)

	repos := repository.NewRepositories(pg)
	uc := usecase.NewUsersUsecase(repos.UserRepository)

	h := handler.NewHandler(uc)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	finalHandler := handler.LoggingMiddleware(handler.APIKeyMiddleware(mux, apiKey))

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      finalHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server started on :8080")
	log.Println("Use X-API-KEY:", apiKey)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func initPostgreConfig() *modules.PostgreConfig {
	return &modules.PostgreConfig{
		Host:        "localhost",
		Port:        "5432",
		Username:    "akmaralomargali",
		Password:    "",
		DBName:      "mydb",
		SSLMode:     "disable",
		ExecTimeout: 5 * time.Second,
	}
}
