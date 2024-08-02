package handlers

import (
	"black-list/internal/models"
	"black-list/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostWorkerHandler(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		POSTWorkerRequest := models.RecordFields{}

		err := json.NewDecoder(r.Body).Decode(&POSTWorkerRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		response, err := s.POSTWorker(context.Background(), POSTWorkerRequest)
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
