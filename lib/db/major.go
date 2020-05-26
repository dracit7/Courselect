package db

// Major information table.
type Major struct {
	ID   int
	Name string
}

// TableName sets the corresponding table name of struct.
func (s Major) TableName() string {
	return "major"
}
