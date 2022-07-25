package repository

import (
	"context"

	"github.com/DarioRoman01/grpc/models"
)

// Respository interface
type Respository interface {
	// GetStudent returns a student with the given id
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	// SetStudent saves a student in the repository
	SetStudent(ctx context.Context, student *models.Student) error
	// GetTest returns a test with the given id
	GetTest(ctx context.Context, id string) (*models.Test, error)
	// SetTest saves a test in the repository
	SetTest(ctx context.Context, test *models.Test) error
	// SetQuestion saves a question in the repository
	SetQuestion(ctx context.Context, question *models.Question) error
	// SetEnrollment saves an enrollment in the repository
	SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	// GetStudentEnrollments returns all students enrollments in the given test
	GetStudentsPerTest(ctx context.Context, testID string) ([]*models.Student, error)
}

// implementation of the repository
var implementation Respository

// SetImplementation sets the implementation of the repository
func SetRepository(repository Respository) {
	implementation = repository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

func SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enrollment)
}

func GetStudentsPerTest(ctx context.Context, testID string) ([]*models.Student, error) {
	return implementation.GetStudentsPerTest(ctx, testID)
}
