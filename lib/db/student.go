package db

import (
	"fmt"
)

// Student information table.
type Student struct {
	id       string
	name     string
	grade    string
	class    int
	major    int
	password string
	email    string
	phone    string
}

// TableName sets the corresponding table name of struct.
func (s Student) TableName() string {
	return "student"
}

// StudentLogin checks if requested user exists in
// the database and if password is correct. If not,
// it will return an error describing the fault.
func StudentLogin(user string, pw string) error {
	var student Student

	db.Where("id = ?", user).Select("password").First(&student)
	if student == (Student{}) {
		return fmt.Errorf("student ID does not exist")
	}

	if student.password != pw {
		return fmt.Errorf("incorrect password")
	}
	return nil
}
