package db

import (
	"fmt"

	"github.com/dracit7/Courselect/setting"
)

// Faculty information table.
type Faculty struct {
	ID       string
	Name     string
	Password string
	Position string
	Email    string
	Phone    string
}

// TableName sets the corresponding table name of struct.
func (s Faculty) TableName() string {
	return "faculty"
}

// FacultyLogin checks if requested user exists in
// the database and if password is correct. If not,
// it will return an error describing the fault.
func FacultyLogin(user string, pw string) error {
	var faculty Faculty
	var cnt int

	db.Where("id = ?", user).Select("password").First(&faculty).Count(&cnt)
	if cnt == 0 {
		return fmt.Errorf("student ID does not exist")
	}

	if faculty.Password != pw {
		return fmt.Errorf("incorrect password")
	}
	return nil
}

// GetFacultyName grabs a faculty's name from database.
func GetFacultyName(id string) string {
	var faculty Faculty

	db.Where("id = ?", id).Select("name").First(&faculty)
	return faculty.Name
}

// GetFaculty grabs a faculty from database.
func GetFaculty(id string) Faculty {
	var faculty Faculty

	db.Where("id = ?", id).First(&faculty)
	return faculty
}

// GetFacultys return all facultys in a page.
func GetFacultys(page int) []Faculty {
	var facultys []Faculty

	db.Offset(page * setting.UI.Pagesize).
		Limit(setting.UI.Pagesize).
		Find(&facultys)
	return facultys
}

// GetFacultyNum return the number of facultys.
func GetFacultyNum() int {
	var count int

	db.Table("faculty").Count(&count)
	return count
}
