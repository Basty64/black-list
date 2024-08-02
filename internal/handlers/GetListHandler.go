package handlers

import (
	"black-list/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GETListHandler(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		response, err := s.GETListOfWorkers(context.Background())
		if err != nil {
			http.Error(w, "Status Forbidden", http.StatusForbidden)
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
	}
}
