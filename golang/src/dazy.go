/*
Name:		dazy.go
Author:		Daniel Middleton <me@daniel-middleton.com>
Description:	DAn's laZY general helper functions
Usage:		See godoc
*/
package main

// Import package dependancies
import (
	"fmt"
	"os"
	"time"
)

// Define global variables
var (
	progName string = "<REPLACE_NAME>"
)

// Desc: Prints a given message to screen in a uniform format
// Usage: screenOut("STATUS", "message")
func screenOut(status string, message string) {
	timestamp := time.Now().Format("20060102T150405")
	fmt.Printf("[%v][%v][%v]: %v \n", progName, timestamp, status, message)
}

// Desc: Generically handles potential errors
// Usage: checkStatus(err, "Error message", "Success message")
func checkStatus(err error, errorMsg string, successMsg string) {
	if err == nil {
		screenOut("SUCCESS", successMsg)
	} else {
		screenOut("ERROR", errorMsg+" - Exiting...")
		screenOut("ERROR", "Error report:")
		fmt.Println("\n", err, "\n")
		os.Exit(1)
	}
}
