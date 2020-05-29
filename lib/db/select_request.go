package db

import (
	"time"

	"github.com/dracit7/Courselect/setting"
)

// Request information table.
type Request struct {
	ID      int
	Course  int
	Student string
	Time    time.Time
}

// TableName sets the corresponding table name of struct.
func (s Request) TableName() string {
	return "select_request"
}

// SelectCourse adds a course to a student's select list.
func SelectCourse(uid string, cid int) {
	request := &Request{
		Course:  cid,
		Student: uid,
		Time:    time.Now(),
	}

	db.Table("select_request").Create(&request)
}

// UnselectCourse adds a course to a student's select list.
func UnselectCourse(uid string, cid int) {
	db.Table("select_request").
		Where("course = ? and student = ?", cid, uid).
		Delete(Request{})
}

// GetSelectedCourses return all selected courses in a page.
func GetSelectedCourses(id string, page int) []Course {
	var courses []Course

	db.Table("select_request").Where("student = ?", id).
		Joins("join course on select_request.course = course.id").
		Joins("join faculty on course.teacher = faculty.id").
		Offset(page * setting.UI.Pagesize).
		Limit(setting.UI.Pagesize).
		Select("course.*, faculty.name as teacher_name").
		Find(&courses)

	for i, course := range courses {
		db.Table("select_request").
			Where("course = ?", course.ID).
			Count(&courses[i].SelectNum)
	}
	return courses
}

// GetSelectedCourseNum return the number of selected
// courses.
func GetSelectedCourseNum(id string) int {
	var count int

	db.Table("select_request").Where("student = ?", id).Count(&count)
	return count
}
