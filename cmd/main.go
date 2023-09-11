package main

import (
	"go-template/config"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	authHandler "go-template/internal/auth/handler"
	authRepo "go-template/internal/auth/repository/mysql"
	authUsecase "go-template/internal/auth/usecase"
)

func main() {
	cfg := config.LoadConfig()
	db := sqlx.MustConnect(cfg.DBDriver, cfg.DBUrl)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// register handlers
	authRepo := authRepo.NewAuthRepo(db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo)
	authHandler.NewAuthHandler(r, authUsecase, logger)

	http.ListenAndServe(":3000", r)
}
