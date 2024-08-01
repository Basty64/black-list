package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RepoPostgres struct {
	connection *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *RepoPostgres {
	return &RepoPostgres{connection: conn}
}

// POSTWorker CREATE USER
func (repo *RepoPostgres) POSTWorker(ctx context.Context, FullName string, Passport int, IsEmployee bool, ActualCompany string, Reason string) (int, error) {
	var id int
	err := repo.connection.QueryRow(ctx, "INSERT INTO list (fullname, passport, is_employee, actual_company, reason) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		FullName,
		Passport,
		IsEmployee,
		ActualCompany,
		Reason).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}

func (repo *RepoPostgres) Close() {
	repo.connection.Close()
}
