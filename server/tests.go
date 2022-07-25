package server

import (
	"context"
	"io"

	"github.com/DarioRoman01/grpc/models"
	"github.com/DarioRoman01/grpc/repository"
	"github.com/DarioRoman01/grpc/testpb"
)

// TestServer is the implementation of the testpb.TestServiceServer interface
type TestServer struct {
	// repository that stores the tests
	repo repository.Respository
	// unimplemented method
	testpb.UnimplementedTestServiceServer
}

// NewTestServer creates a new TestServer instance with a given repository
func NewTestServer(repo repository.Respository) *TestServer {
	return &TestServer{repo: repo}
}

// function that gives the implementation of the GetTest method
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

// function that gives the implementation of the SetTest method
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

func (s *TestServer) SetQuestion(stream testpb.TestService_SetQuestionServer) error {
	for {
		question, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			return err
		}

		err = s.repo.SetQuestion(stream.Context(), &models.Question{
			ID:       question.GetId(),
			Answer:   question.GetAnswer(),
			Question: question.GetQuestion(),
			TestID:   question.GetTestId(),
		})

		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}
