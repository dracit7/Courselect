package db

import (
	"fmt"

	"github.com/dracit7/Courselect/setting"
)

// Course information table.
type Course struct {
	ID          int
	Name        string
	Teacher     string
	TeacherName string
	Credit      int
	Capacity    int
	Sdate       int
	Edate       int
	Day         string
	Stime       string
	Etime       string
	Valid       string
}

// TableName sets the corresponding table name of struct.
func (s Course) TableName() string {
	return "course"
}

// GetCourseByID returns a course with certain ID.
func GetCourseByID(id int) (*Course, error) {
	var course Course
	var cnt int

	db.Select("*").Find(&course).Count(&cnt)
	if cnt == 0 {
		return nil, fmt.Errorf("no such id")
	}
	return &course, nil
}

// GetSelectableCourses return all selectable courses in a page.
func GetSelectableCourses(page int) []Course {
	var courses []Course

	db.Where("valid = ?", 1).
		Joins("join faculty on course.teacher = faculty.id").
		Offset(page * setting.UI.Pagesize).
		Limit(setting.UI.Pagesize).
		Select("course.*, faculty.name as teacher_name").
		Find(&courses)
	return courses
}

// GetSelectableCourseNum return the number of selectable
// courses.
func GetSelectableCourseNum() int {
	var count int

	db.Table("course").Where("valid = ?", 1).Count(&count)
	return count
}

// GetTeachingCourses return all courses submitted by target
// teacher in a page.
func GetTeachingCourses(teacher string, page int) []Course {
	var courses []Course

	db.Where("teacher = ?", teacher).
		Offset(page * setting.UI.Pagesize).
		Limit(setting.UI.Pagesize).
		Select("*").Find(&courses)
	return courses
}
