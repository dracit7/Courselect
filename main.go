package main

import (
	"fmt"
	"os"

	"github.com/dracit7/Courselect/lib/db"
	"github.com/dracit7/Courselect/lib/log"
	"github.com/dracit7/Courselect/router"
	"github.com/dracit7/Courselect/setting"

	"github.com/fvbock/endless"
	"github.com/urfave/cli"
)

func main() {

	var configFile string

	// Handle command-line arguments.
	app := cli.NewApp()
	app.Name = "Courselect"
	app.Usage = "Courselect is an online platform for publishing and selecting courses."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "conf",
			Value:       "conf/debug.conf",
			Usage:       "specify config file",
			Destination: &configFile,
		},
	}
	app.Action = func(c *cli.Context) error {
		Start(configFile)
		return nil
	}

	// Run the application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Start the application
func Start(configFile string) {

	// Set up each module.
	setting.Setup(configFile)
	log.Setup(os.Stdout)
	db.Setup()

	router := router.Setup()
	endpoint := fmt.Sprintf(":%d", setting.Server.Port)

	server := endless.NewServer(endpoint, router)
	server.BeforeBegin = func(add string) {
		log.Info("*** Started the server ***")
	}

	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server has been shutted down.")
	}
}
