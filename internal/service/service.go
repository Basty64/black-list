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

func (s *Service) AddUser(login string, passwordHash string) (string, error) {

	return "", nil
}

func (s *Service) POSTWorker(ctx context.Context, rf models.RecordFields) (int, error) {
	return s.repoPostgres.POSTWorker(ctx, rf.FullName, rf.Passport, rf.IsEmployee, rf.ActualCompany, rf.Reason)
}
