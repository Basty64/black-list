package handlers

import (
	"black-list/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GETWorkerHandler(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, _ := strconv.Atoi(r.PathValue("id"))
		response, err := s.GETWorker(context.Background(), id)
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
