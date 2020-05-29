package db

import (
	"fmt"

	"github.com/dracit7/Courselect/setting"
)

// Student information table.
type Student struct {
	ID        string
	Name      string
	Grade     string
	Class     int
	Major     int
	MajorName string
	Password  string
	Email     string
	Phone     string
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
	var cnt int

	db.Where("id = ?", user).Select("password").First(&student).Count(&cnt)
	if cnt == 0 {
		return fmt.Errorf("student ID does not exist")
	}

	if student.Password != pw {
		return fmt.Errorf("incorrect password")
	}
	return nil
}

// GetStudentName grabs a student's name from database.
func GetStudentName(id string) string {
	var student Student

	db.Where("id = ?", id).Select("name").First(&student)
	return student.Name
}

// GetStudent grabs a student from database.
func GetStudent(id string) Student {
	var student Student

	db.Joins("join major on major.id = student.major").
		Where("student.id = ?", id).
		Select("student.*, major.name as major_name").
		First(&student)
	return student
}

// GetStudents return all students in a page.
func GetStudents(page int) []Student {
	var students []Student

	db.Joins("join major on major.id = student.major").
		Offset(page * setting.UI.Pagesize).
		Limit(setting.UI.Pagesize).
		Select("student.*, major.name as major_name").
		Find(&students)
	return students
}

// GetStudentNum return the number of students.
func GetStudentNum() int {
	var count int

	db.Table("student").Count(&count)
	return count
}

// GetStudentsInCourse return all students in a page.
func GetStudentsInCourse(cid int, page int) []Student {
	var students []Student

	db.Joins("join major on major.id = student.major").
		Joins("join select_request on select_request.student = student.id").
		Where("select_request.course = ?", cid).
		Offset(page * setting.UI.Pagesize).
		Limit(setting.UI.Pagesize).
		Select("student.*, major.name as major_name").
		Find(&students)
	return students
}

// GetStudentInCourseNum return the number of students.
func GetStudentInCourseNum(cid int) int {
	var count int

	db.Table("student").
		Joins("join select_request on select_request.student = student.id").
		Where("select_request.course = ?", cid).
		Count(&count)
	return count
}
