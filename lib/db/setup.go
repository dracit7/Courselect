package db

import (
	"fmt"

	"github.com/dracit7/Courselect/lib/log"
	"github.com/dracit7/Courselect/setting"

	"github.com/jinzhu/gorm"

	// Blank import sql drivers
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// Setup sets up the connection to database.
func Setup() {
	var err error

	db, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DB.User, setting.DB.Password, setting.DB.Database,
	))
	if err != nil {
		log.Fatal(err.Error())
	}

}
