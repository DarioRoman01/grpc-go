package repository

import (
	"context"

	"github.com/DarioRoman01/grpc/models"
)

type Respository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var implementation Respository

func SetRepository(repository Respository) {
	implementation = repository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}
