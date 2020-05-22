package db

import (
	"fmt"
	"time"

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
	Sdate       time.Time
	Edate       time.Time
	Day         string
	Stime       time.Time
	Etime       time.Time
	Valid       int
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
