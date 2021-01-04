package logger

import (
	"log"
	"os"
	"strings"
	"time"
)

func (logg AccessLog) Send() {
	if logg.source == "" {
		logg.source = "Unavailable"
	}
	if logg.location == "" {
		logg.location = "Unavailable"
	}
	if !logg.showAuth {
		logg.authUser = ""
	}
	//If debugging is enabled on the logger, send this log to the terminal also
	if logg.debug {
		var entry strings.Builder
		present := time.Now()

		entry.WriteString("\x1b[34m[  ACCESS  ]")
		entry.WriteString(present.Format("[ 15:04:05 ] "))
		entry.WriteString(logg.source + ": ")
		entry.WriteString("\x1b[36m" + logg.method + "\x1b[34m -> ")
		entry.WriteString(logg.location + " <- ")
		entry.WriteString(logg.authUser)

		_, err := os.Stdout.WriteString(entry.String() + "\x1b[0m\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	//Log the access message to file
	var entry strings.Builder
	present := time.Now()

	entry.WriteString(present.Format("15:04:05-"))
	entry.WriteString(logg.source + "-")
	entry.WriteString(logg.method + "-")
	entry.WriteString(logg.location + "-")
	entry.WriteString(logg.authUser + "-\n")

	logFile, err := os.OpenFile(logg.accessPath+present.Format("02-01-2006.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		defer logFile.Close()
	}

	_, err = logFile.WriteString(entry.String())
	if err != nil {
		log.Fatal(err)
	}
}
