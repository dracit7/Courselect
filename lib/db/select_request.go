package db

// Request information table.
type Request struct {
	ID       string
	Name     string
	Password string
	Position int
}

// TableName sets the corresponding table name of struct.
func (s Request) TableName() string {
	return "select_request"
}
