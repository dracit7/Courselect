package db

// Major information table.
type Major struct {
	ID       string
	Name     string
	Password string
	Position int
}

// TableName sets the corresponding table name of struct.
func (s Major) TableName() string {
	return "major"
}
