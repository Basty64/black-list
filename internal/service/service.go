package service

import (
	"black-list/internal/models"
	"black-list/internal/repository/migrations/postgres"
	"context"
)

//type ServiceMethods interface {
//	AddWorker(login string, passwordHash string) (string, error)
//	GetList() ([]models.RecordFields, error)
//	GetWorker(id int) (models.RecordFields, error)
//	POSTWorker(context.Context, models.RecordFields) (int, error)
//}

type Service struct {
	ctx          context.Context
	repoPostgres *postgres.RepoPostgres
}

func NewService(repo *postgres.RepoPostgres) *Service {
	return &Service{repoPostgres: repo}
}

//TO DO:
//func (s *Service) AddUser(login string, passwordHash string) (string, error) {
//	return "", nil
//}

func (s *Service) POSTWorker(ctx context.Context, rf models.RecordFields) (int, error) {
	return s.repoPostgres.POSTWorker(ctx, rf.FullName, rf.Passport, rf.IsEmployee, rf.ActualCompany, rf.Reason)
}

func (s *Service) GETListOfWorkers(ctx context.Context) ([]models.Worker, error) {
	return s.repoPostgres.GETListOfWorkers(ctx)
}

func (s *Service) GETWorker(ctx context.Context, ID int) (models.Worker, error) {
	return s.repoPostgres.GETWorker(ctx, ID)
}

func (s *Service) UpdateWorker(ctx context.Context, worker models.Worker, ID int) error {
	return s.repoPostgres.UpdateWorker(ctx, worker.FullName, worker.Passport, worker.IsEmployee, worker.ActualCompany, worker.Reason, ID)
}

func (s *Service) DeleteWorker(ctx context.Context, ID int) error {
	return s.repoPostgres.DeleteWorker(ctx, ID)
}
