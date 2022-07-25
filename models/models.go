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
