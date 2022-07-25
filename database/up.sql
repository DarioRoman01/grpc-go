DROP TABLE IF EXISTS students;

CREATE TABLE students (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  age INTEGER NOT NULL
);

DROP TABLE IF EXISTS tests;

CREATE TABLE tests (
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS questions;

CREATE TABLE questions (
  id VARCHAR(32) PRIMARY KEY,
  test_id VARCHAR(32) NOT NULL,
  question VARCHAR(255) NOT NULL,
  answer VARCHAR(255) NOT NULL,
  FOREIGN KEY (test_id) REFERENCES tests(id)
);

DROP TABLE IF EXISTS enrollments;

CREATE TABLE enrollments (
  student_id VARCHAR(32) NOT NULL,
  test_id VARCHAR(32) NOT NULL,
  FOREIGN KEY (student_id) REFERENCES students(id),
  FOREIGN KEY (test_id) REFERENCES tests(id)
);

DROP TABLE IF EXISTS answers;

CREATE TABLE answers (
  student_id VARCHAR(32) NOT NULL,
  question_id VARCHAR(32) NOT NULL,
  test_id VARCHAR(32) NOT NULL,
  answer VARCHAR(255) NOT NULL,
  correct_answer VARCHAR(255) NOT NULL,
  correct BOOLEAN NOT NULL,
  FOREIGN KEY (student_id) REFERENCES students(id),
  FOREIGN KEY (question_id) REFERENCES questions(id),
  FOREIGN KEY (test_id) REFERENCES tests(id)
);

INSERT INTO students (id, name, age) VALUES ('s1', 'John', 20);
INSERT INTO students (id, name, age) VALUES ('s2', 'Jane', 21);
INSERT INTO students (id, name, age) VALUES ('s3', 'Jack', 22);

INSERT INTO tests (id, name) VALUES ('t1', 'GOLANG 101');

INSERT INTO questions (id, test_id, question, answer) VALUES ('q1', 't1', 'What is the name of the language?', 'GOLANG');
INSERT INTO questions (id, test_id, question, answer) VALUES ('q2', 't1', 'What is the name of the company?', 'Google');
INSERT INTO questions (id, test_id, question, answer) VALUES ('q3', 't1', 'What is the name of the framework?', 'Go');