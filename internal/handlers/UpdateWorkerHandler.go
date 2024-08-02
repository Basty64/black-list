package handlers

import (
	"black-list/internal/models"
	"black-list/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateWorkerHandler(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		UpdateWorkerRequest := models.Worker{}

		id, _ := strconv.Atoi(r.PathValue("id"))
		err := json.NewDecoder(r.Body).Decode(&UpdateWorkerRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		var response string
		err = s.UpdateWorker(context.Background(), UpdateWorkerRequest, id)
		if err != nil {
			http.Error(w, "Status Forbidden", http.StatusForbidden)
			fmt.Println(err)
			response = "ERROR"
		}

		w.Header().Set("Content-Type", "application/json")

		response = fmt.Sprintf("Worker %d updated successfully", id)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
	}
}
