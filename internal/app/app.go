package app

import (
	"black-list/internal/handlers"
	"black-list/internal/logs"
	"black-list/internal/repository/migrations/postgres"
	"black-list/internal/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type App struct {
	mux  *http.ServeMux
	err  error
	repo *postgres.RepoPostgres
}

func New() (*App, error) {

	return &App{}, nil

}

func (a *App) Router(s *service.Service) http.Handler {

	mux := http.NewServeMux()

	//authorization
	mux.Handle("POST /api/v0/registration", logs.RequestLogger(handlers.UserRegistrationHandler(s)))
	mux.Handle("GET /api/v0/authentication", logs.RequestLogger(handlers.UserAuthenticationHandler(s)))

	//admin role
	mux.Handle("GET /api/v0/list", logs.RequestLogger(handlers.GETListHandler(s)))
	mux.Handle("GET /api/v0/list/{id}", logs.RequestLogger(handlers.GETWorkerHandler(s)))
	mux.Handle("POST /api/v0/list", logs.RequestLogger(handlers.PostWorkerHandler(s)))
	mux.Handle("DELETE /api/v0/list/{id}", logs.RequestLogger(handlers.DeleteWorkerHandler(s)))
	mux.Handle("PATCH /api/v0/list/{id}", logs.RequestLogger(handlers.UpdateWorkerHandler(s)))
	mux.Handle("/api/v0/test-frontend", logs.RequestLogger(handlers.FrontendHandler()))

	a.mux = mux
	return a.mux
}

func (a *App) Run() error {

	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	postgresDB := os.Getenv("DATABASE_URL_POSTGRES")

	poolConfig, err := pgxpool.ParseConfig(postgresDB)
	if err != nil {
		log.Fatalf("failed to parse the config: %w", err)
	}

	dbConn, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %w", err)
	}

	repository := postgres.NewRepository(dbConn)

	defer repository.Close()

	serviceStruct := service.NewService(repository)

	server := &http.Server{Addr: host + ":" + port, Handler: a.Router(serviceStruct)}

	fmt.Printf("Сервер запущен на порту %s и успешно работает\n", port)
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil

}
