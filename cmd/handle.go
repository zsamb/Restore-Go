package cmd

import (
	"fmt"
	"log"
	"os"
)

//Handler Handles command arguments
func Handler() {
	//Check argument is passed
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "help":
			fmt.Println("The following commands are available:")
			fmt.Println("'help' Shows a list of commands")
			fmt.Println("'start' Begins Restore if its not already running")
			fmt.Println("'status' Displays Restore status")
		case "start":
			Start()
		default:
			log.Fatalf("Invalid command '%s'\n- Try 'help' for a list of commands", os.Args[1])
		}
	} else {
		log.Fatal("Please specify a command. \n- Try 'help' for a list of commands")
	}
}
