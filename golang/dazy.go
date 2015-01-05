/*
Name:		dazy.go
Author:		Daniel Middleton <me@daniel-middleton.com>
Description:	DAn's laZY general helper functions
Usage:		See function comments
Notes: 		Assumes progName string is defined elsewhere
*/
package main

// Import package dependancies
import "log"

// Define global variables
var ()

// Desc: Prints a given message to screen and log in a uniform format
// Usage: logOut("STATUS", "message")
func logOut(status string, message string) {
	switch status {
	case "INFO":
		log.Println(status, message)
	case "SUCCESS":
		log.Println(status, message)
	case "ERROR":
		log.Panicln(status, message)
	default:
		log.Panicln("Unhandled logging status (%s). Exiting...", status)
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
