/*
Name:		dazy.go
Author:		Daniel Middleton <me@daniel-middleton.com>
Description:	DAn's laZY general helper functions
Usage:		See function comments
Notes: 		Assumes progName string is defined elsewhere
*/
package main

// Import package dependancies
import (
	"fmt"
	"log"
	"os"
)

// Define global variables
var ()

// Desc: Prints a given message to screen and log in a uniform format
// Usage: logOut("STATUS", "message")
func logOut(status string, message string) {
	f, err := os.Create(fmt.Sprintf("/tmp/%s.log", progName))
	if err == nil {
		log.Fatalf(progName, "ERROR", "Could not create log. Exiting...")
	}

	switch status {
	case "INFO":
		log.Println(progName, status, message)
	case "SUCCESS":
		log.Println(progName, status, message)
	case "ERROR":
		log.Fatalf(progName, status, message)
	default:
		log.Panicln(progName, "Unhandled logging status (%s). Exiting...", status)
	}
}

// Desc: Generically handles errors
// Usage: checkStatus(err, "Error message", "Success message")
func checkStatus(err error, errorMsg string, successMsg string) {
	if err == nil {
		logOut("SUCCESS", successMsg)
	} else {
		logOut("ERROR", errorMsg+" - Exiting...")
	}
}
