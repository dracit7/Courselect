package db

import (
	"github.com/dracit7/Courselect/lib/log"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Setup sets up the connection to database.
func Setup() {
	db, err := gorm.Open("mysql", "@localhost")
	if err != nil {
		log.Fatal("Failed to connect database.")
	}
	defer db.Close()
}
