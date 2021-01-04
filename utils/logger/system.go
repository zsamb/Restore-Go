package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func (logg SystemLog) Send() {
	var entry strings.Builder
	present := time.Now()

	//Clean the message argument to prevent problems when reading logs
	logg.message = strings.ReplaceAll(logg.message, "-", "")

	if logg.error {
		entry.WriteString("\x1b[31m")
	}
	entry.WriteString("[  SYSTEM  ]")
	entry.WriteString(present.Format("[ 15:04:05 ] "))
	entry.WriteString(logg.message + "\x1b[0m\n")

	_, err := os.Stdout.WriteString(entry.String())
	if err != nil {
		log.Fatal(err)
	}

	var fileEntry strings.Builder
	fileEntry.WriteString(present.Format("15:04:05-"))
	fileEntry.WriteString(fmt.Sprint(logg.error) + "-")
	fileEntry.WriteString(logg.message + "-\n")

	logFile, err := os.OpenFile(logg.systemPath+present.Format("02-01-2006.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		defer logFile.Close()
	}

	_, err = logFile.WriteString(fileEntry.String())
	if err != nil {
		log.Fatal(err)
	}
}
