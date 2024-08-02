package handlers

import (
	"black-list/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func DeleteWorkerHandler(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, _ := strconv.Atoi(r.PathValue("id"))

		var response string
		err := s.DeleteWorker(context.Background(), id)
		if err != nil {
			http.Error(w, "Status Forbidden", http.StatusForbidden)
			fmt.Println(err)
			response = "ERROR"
		}

		w.Header().Set("Content-Type", "application/json")
		now := time.Now()
		response = fmt.Sprintf("Worker %d deleted successfully at %s", id, now.Format(time.RFC3339))
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
	}
}
