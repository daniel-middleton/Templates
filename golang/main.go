/*
Name:		<REPLACE_NAME>
Author:		Daniel Middleton <me@daniel-middleton.com>
Description:	<REPLACE_DESC>
Usage:		See ./<REPLACE_NAME> -h or --help
Notes: 		Assumes dazy.go exists in package for helper functions
*/
package main

// Import package dependancies
import (
	"fmt"
	"os"
)

// Define global variables
const (
	progName string = "<REPLACE_NAME>" // Required by dazy.go
)

// Desc:
// Usage:
func newFunc() {
	logOut("INFO", "Started newFunc...")

	// Do stuff
	//	checkStatus(err,
	//		"Error message.",
	//		"Success message.")

	logOut("INFO", "Finished newFunc.")
}

// Desc: Auto-runs on program execution
// Usage: n/a
func init() {

}

// Desc: Auto-runs after init function
// Usage: n/a
func main() {
	fmt.Println()
	logOut("INFO", "Script started...")

	// Execute functions

	logOut("INFO", "Script ended.")
	fmt.Println()
	os.Exit(0)
}
