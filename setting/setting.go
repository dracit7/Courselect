package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var config *ini.File

// ServerConfig stores configurations of the http server.
type ServerConfig struct {
	Mode string
	Port int
}

// DBConfig stores configurations of the database.
type DBConfig struct {
	User     string
	Password string
	Database string
}

// Server stores configurations of the server.
var Server = &ServerConfig{}

// DB stores configurations of the database.
var DB = &DBConfig{}

// Setup reads all settings from the config file and
// load them into structures.
func Setup(configFile string) {
	var err error

	config, err = ini.Load(configFile)
	if err != nil {
		log.Fatalf("Failed to open the config file '%s'.", configFile)
	}

	load("server", Server)
	load("db", DB)
}

// load loads a section from the config file.
func load(section string, target interface{}) {

	err := config.Section(section).MapTo(target)
	if err != nil {
		log.Println("Failed to load section " + section + " from config file.")
		log.Fatalln("Error: " + err.Error())
	}
}
