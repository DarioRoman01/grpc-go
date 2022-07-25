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

// function that gives the implementation of the GetStudent method
func (r *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	var student models.Student
	err := r.db.QueryRowContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id).Scan(&student.ID, &student.Name, &student.Age)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// function that gives the implementation of the SetStudent method
func (r *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.ID, student.Name, student.Age)
	return err
}

// function that gives the implementation of the GetTest method
func (r *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	var test models.Test
	err := r.db.QueryRowContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id).Scan(&test.ID, &test.Name)
	if err != nil {
		return nil, err
	}

	return &test, nil
}

// function that gives the implementation of the SetTest method
func (r *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.ID, test.Name)
	return err
}

// function that gives the implementation of the SetQuestion method
func (r *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO questions (id, question, answer, test_id) VALUES ($1, $2, $3, $4)", question.ID, question.Question, question.Answer, question.TestID)
	return err
}

// function that gives the implementation of the SetEnrollment method
func (r *PostgresRepository) SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO enrollments (student_id, test_id) VALUES ($1, $2)", enrollment.StudentID, enrollment.TestID)
	return err
}

// function that gives the implementation of the GetStudentsPerTest method
func (r *PostgresRepository) GetStudentsPerTest(ctx context.Context, testID string) ([]*models.Student, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id IN (SELECT student_id FROM enrollments WHERE test_id = $1)", testID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var students []*models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	return students, nil
}
