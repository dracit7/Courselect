package db

// Result information table.
type Result struct {
	ID      int
	Course  int
	Student string
	Grade   int
}

// TableName sets the corresponding table name of struct.
func (s Result) TableName() string {
	return "select_result"
}
