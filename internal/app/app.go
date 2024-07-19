package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type App struct {
	mux *http.ServeMux
	err error
}

func New() (*App, error) {

	return &App{}, nil

}

func (a *App) Router() http.Handler {

	mux := http.NewServeMux()

	//authorization
	mux.HandleFunc("POST /api/v0/registration", nil)
	mux.HandleFunc("GET /api/v0/authentication", nil)

	//admin role
	mux.HandleFunc("GET /api/v0/list", nil)
	mux.HandleFunc("GET /api/v0/list/{id}", nil)
	mux.HandleFunc("DELETE /api/v0/list/{id}", nil)
	mux.HandleFunc("PATCH /api/v0/list/{id}", nil)
	mux.HandleFunc("POST /api/v0/list/{id}", nil)

	a.mux = mux
	return a.mux
}

func (a *App) Run() error {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")

	server := &http.Server{Addr: host + ":" + port, Handler: a.Router()}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil

}
