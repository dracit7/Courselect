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
