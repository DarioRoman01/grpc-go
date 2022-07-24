package database

import (
	"context"
	"database/sql"

	"github.com/DarioRoman01/grpc/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

// function that creates a new PostgresRepository instance with a given url to connect to the database
func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (r *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student
	err := r.db.QueryRowContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id).Scan(&student.ID, &student.Name, &student.Age)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.ID, student.Name, student.Age)
	return err
}

func (r *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	var test models.Test
	err := r.db.QueryRowContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id).Scan(&test.ID, &test.Name)
	if err != nil {
		return nil, err
	}

	return &test, nil
}

func (r *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.ID, test.Name)
	return err
}
