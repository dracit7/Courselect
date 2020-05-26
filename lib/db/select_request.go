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
