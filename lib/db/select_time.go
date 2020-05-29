package db

import (
	"time"
)

// Time information table.
type Time struct {
	ID    int
	Major int
	Stime time.Time
	Etime time.Time
}

// TableName sets the corresponding table name of struct.
func (s Time) TableName() string {
	return "select_time"
}

// GetTimeRange return the time range of course selection.
func GetTimeRange(uid string) Time {
	var time Time

	db.Table("select_time").
		Joins("join student on student.major = select_time.major").
		Where("student.id = ?", uid).
		Select("select_time.*").
		First(&time)

	return time
}

// GetApplyDeadline returns the deadline of creating courses.
func GetApplyDeadline() time.Time {
	var time Time

	db.Table("select_time").
		Order("stime asc").First(&time)
	return time.Stime
}

// GetResultTime returns the time of generating result.
func GetResultTime() time.Time {
	var time Time

	db.Table("select_time").
		Order("etime desc").First(&time)
	return time.Etime
}
