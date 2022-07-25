package models

// representation of a student
type Student struct {
	// id of the student
	ID string `json:"id"`
	// name of the student
	Name string `json:"name"`
	// age of the student
	Age int32 `json:"age"`
}

// representation of a test
type Test struct {
	// id of the test
	ID string `json:"id"`
	// name of the test
	Name string `json:"name"`
}

// representation of test question
type Question struct {
	// id of the question
	ID string `json:"id"`
	// question of the question
	Question string `json:"question"`
	// answer of the question
	Answer string `json:"answer"`
	// test id where the question belongs to
	TestID string `json:"test_id"`
}

// represents a enrollment of a student in a test
type Enrollment struct {
	// student id of the enrollment
	StudentID string `json:"student_id"`
	// test id of the enrollment
	TestID string `json:"test_id"`
}

// representation of a student response to a question
type Answer struct {
	// student id of the enrollment
	StudentID string `json:"student_id"`
	// test id of the enrollment
	TestID string `json:"test_id"`
	// question id of the answer
	QuestionID string `json:"question_id"`
	// answer of the question
	Answer string `json:"answer"`
	// correct answer of the question
	CorrectAnswer string `json:"correct_answer"`
	// the answer is correct or not
	Correct bool `json:"correct"`
}

type StudentScore struct {
	// student id of the enrollment
	StudentID string `json:"student_id"`
	// test id of the enrollment
	TestID string `json:"test_id"`
	// score of the student
	Score int32 `json:"score"`
	// total number of questions in the test
	Total int32 `json:"total"`
}
