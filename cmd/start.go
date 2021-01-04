package cmd

import (
	"github.com/Samb8104/Restore/utils/config"
	"github.com/Samb8104/Restore/utils/database"
	"github.com/Samb8104/Restore/utils/logger"
	"log"
)

//Start Restore
func Start() {
	//Validate and initiate logger
	logg := logger.Create("logs/access/", "logs/activity/", "logs/system/", false)
	err := logg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	//Read & validate config
	logg.System("Reading configuration files..", false)
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	logg.System("Updating logger information..", false)
	logg.Debug = conf.Debug

	//Validate database and setup tables if required
	logg.System("Testing connection to database..", false)
	db, err := database.Initialise(*conf)
	if err != nil {
		log.Fatal(err)
	}
	logg.System("Validated connection to "+db.Host(), false)

	//First run tasks

	//Start webserver
	//logg.System(fmt.Sprintf("Beginning webserver at :%d", conf.Web.Port), false)
	//web := webserver.Initialise(conf.Web.Port)

	//Start automater

	//logg.Access("POST", "81.64.32.75", "/panel/dash", "Samb8104", true)
	//logg.System("Encountered an error receiving status data.", true)
	//logg.Activity("backup", "Creating new backup.")
	select {}
}
