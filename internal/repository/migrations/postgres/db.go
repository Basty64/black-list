package postgres

import (
	"black-list/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
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

// GETListOfWorkers SHOW LIST OF WORKERS
func (repo *RepoPostgres) GETListOfWorkers(ctx context.Context) ([]models.Worker, error) {
	var ArrayOfWorkers []models.Worker
	rows, err := repo.connection.Query(ctx, "SELECT * FROM LIST")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row models.Worker
		err := rows.Scan(
			&row.ID,
			&row.FullName,
			&row.Passport,
			&row.IsEmployee,
			&row.ActualCompany,
			&row.Reason,
			&row.AddedAt,
			&row.UpdatedAt)
		if err != nil {
			return nil, err
		}
		ArrayOfWorkers = append(ArrayOfWorkers, row)
	}
	return ArrayOfWorkers, nil
}

func (repo *RepoPostgres) GETWorker(ctx context.Context, id int) (models.Worker, error) {
	row, err := repo.connection.Query(ctx, "SELECT * FROM LIST WHERE id = $1", id)
	if err != nil {
		return models.Worker{}, err
	}
	defer row.Close()
	var worker models.Worker
	for row.Next() {
		err := row.Scan(
			&worker.ID,
			&worker.FullName,
			&worker.Passport,
			&worker.IsEmployee,
			&worker.ActualCompany,
			&worker.Reason,
			&worker.AddedAt,
			&worker.UpdatedAt)
		if err != nil {
			return models.Worker{}, err
		}
	}
	return worker, nil
}

func (repo *RepoPostgres) UpdateWorker(ctx context.Context, FullName string, Passport int, IsEmployee bool, ActualCompany string, Reason string, ID int) error {

	now := time.Now()
	_, err := repo.connection.Exec(ctx, "UPDATE list SET fullname = $1, passport = $2, is_employee = $3, actual_company = $4, reason = $5, update_at = $6 WHERE id = $7",
		FullName, Passport, IsEmployee, ActualCompany, Reason, now, ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (repo *RepoPostgres) DeleteWorker(ctx context.Context, ID int) error {

	_, err := repo.connection.Exec(ctx, "DELETE FROM list WHERE id = $1", ID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (repo *RepoPostgres) Close() {
	repo.connection.Close()
}

//func (repo *RepoPostgres) ValidateDuplicates(ctx context.Context, Passport int) error {

// Если хотим узнать количество
//	count := 0
//
//	row := repo.connection.QueryRow(ctx, "select count(fullname) from list where passport = &1", Passport)
//	err := row.Scan(&count)
//	if err != nil {
//		log.Fatal(err)
//	}
//	if count > 0 {
//		err.Error()
//	}
//	return nil
//}
// Если хотим узнать наличие работника
// var exists bool
//    err = conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM list WHERE passport = $1 OR email = $2)", Passport, email).Scan(&exists)
//    if err != nil {
//        log.Fatal(err)
//    }
//    if exists {
//        fmt.Println("Ошибка: Нарушение уникальности. Username или Email уже существует.")
//        return
//    }
