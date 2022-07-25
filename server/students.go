package server

import (
	"context"

	"github.com/DarioRoman01/grpc/models"
	"github.com/DarioRoman01/grpc/repository"
	"github.com/DarioRoman01/grpc/studentpb"
)

// StudentServer is the implementation of the studentpb.StudentServiceServer interface
type StudentsServer struct {
	// repository that stores the students
	repo repository.Respository
	// unimplemented method
	studentpb.UnimplementedStudentServiceServer
}

// function that creates a new Server instance with a given repository
func NewStudentsServer(repo repository.Respository) *StudentsServer {
	return &StudentsServer{repo: repo}
}

// function that gives the implementation of the GetStudent method
func (s *StudentsServer) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &studentpb.Student{
		Id:   student.ID,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

// function that gives the implementation of the SetStudent method
func (s *StudentsServer) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	err := s.repo.SetStudent(ctx, &models.Student{
		ID:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	})
	if err != nil {
		return nil, err
	}

	return &studentpb.SetStudentResponse{Id: req.GetId()}, nil
}
