package server

import (
	"context"

	"github.com/DarioRoman01/grpc/models"
	"github.com/DarioRoman01/grpc/repository"
	"github.com/DarioRoman01/grpc/studentpb"
)

type Server struct {
	repo repository.Respository
	studentpb.UnimplementedStudentServiceServer
}

// function that creates a new Server instance with a given repository
func NewServer(repo repository.Respository) *Server {
	return &Server{repo: repo}
}

// function that gives the implementation of the GetStudent method
func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
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
func (s *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
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
