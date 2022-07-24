package server

import (
	"context"

	"github.com/DarioRoman01/grpc/models"
	"github.com/DarioRoman01/grpc/repository"
	"github.com/DarioRoman01/grpc/testpb"
)

type TestServer struct {
	repo repository.Respository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Respository) *TestServer {
	return &TestServer{repo: repo}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &testpb.Test{
		Id:   test.ID,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SeTestResponse, error) {
	err := s.repo.SetTest(ctx, &models.Test{
		ID:   req.GetId(),
		Name: req.GetName(),
	})
	if err != nil {
		return nil, err
	}

	return &testpb.SeTestResponse{Id: req.GetId()}, nil
}
