package repository

import "black-list/internal/models"

type Repository interface {
	POSTWorker(user models.User) (int, error)
	GETWorker(id int) (models.User, error)
	GETAllWorkers() ([]models.User, error)
	DELETEWorker(id int) error
	UPDATEWorker(user models.User) (int, error)
}
