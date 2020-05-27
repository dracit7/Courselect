package db

import (
	"time"
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
