package server

import (
	"context"
	"io"

	"github.com/DarioRoman01/grpc/models"
	"github.com/DarioRoman01/grpc/repository"
	"github.com/DarioRoman01/grpc/studentpb"
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

func (s *TestServer) EnrollStudent(stream testpb.TestService_EnrollStudentServer) error {
	for {
		enrollment, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			return err
		}

		err = s.repo.SetEnrollment(stream.Context(), &models.Enrollment{
			StudentID: enrollment.GetStudentId(),
			TestID:    enrollment.GetTestId(),
		})

		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) GetStudentPerTest(req *testpb.GetStudentPerTestRequest, stream testpb.TestService_GetStudentPerTestServer) error {
	enrollment, err := s.repo.GetStudentsPerTest(context.Background(), req.GetTestId())
	if err != nil {
		return err
	}

	for _, e := range enrollment {
		err := stream.Send(&studentpb.Student{
			Id:   e.ID,
			Name: e.Name,
			Age:  e.Age,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TestServer) TakeTest(stream testpb.TestService_TakeTestServer) error {
	for {
		msg, err := stream.Recv()
		testId := msg.GetTestId()
		if err == io.EOF {
			return nil
		}

		questions, err := s.repo.GetQuestionsPerTest(context.Background(), testId)
		if err != nil {
			return err
		}

		i := 0
		var currentQuestion *models.Question
		for {
			if i >= len(questions) {
				return nil
			}

			currentQuestion = questions[i]

			err := stream.Send(&testpb.Question{
				Id:       currentQuestion.ID,
				Question: currentQuestion.Question,
			})

			if err != nil {
				return err
			}

			i++

			answer, err := stream.Recv()
			if err == io.EOF {
				return nil
			}

			if err != nil {
				return err
			}

			err = s.repo.SetAnswer(context.Background(), &models.Answer{
				StudentID:     msg.GetStudentId(),
				TestID:        testId,
				QuestionID:    currentQuestion.ID,
				Answer:        answer.GetAnswer(),
				CorrectAnswer: currentQuestion.Answer,
				Correct:       answer.GetAnswer() == currentQuestion.Answer,
			})

			if err != nil {
				return err
			}
		}

	}
}

func (s *TestServer) GetStudentAnswer(stream testpb.TestService_GetStudentAnswerServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		answers, err := s.repo.GetStudentAnswers(context.Background(), msg.GetStudentId(), msg.GetTestId())
		if err != nil {
			return err
		}

		for _, a := range answers {
			err := stream.Send(&testpb.AnswerResponse{
				StudentId:     a.StudentID,
				TestId:        a.TestID,
				QuestionId:    a.QuestionID,
				Answer:        a.Answer,
				CorrectAnswer: a.CorrectAnswer,
				Correct:       a.Correct,
			})

			if err != nil {
				return err
			}
		}
	}
}

func (s *TestServer) GetStudentScore(stream testpb.TestService_GetStudentScoreServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		score, err := s.repo.GetStudentScore(context.Background(), msg.GetStudentId(), msg.GetTestId())
		if err != nil {
			return err
		}

		err = stream.Send(&testpb.ScoreResponse{
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
			Score:     score.Score,
			Total:     score.Total,
		})

		if err != nil {
			return err
		}
	}
}
