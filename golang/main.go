/*
Name:		<REPLACE_NAME>
Author:		Daniel Middleton <me@daniel-middleton.com>
Description:	<REPLACE_DESC>
Usage:		See ./<REPLACE_NAME> -h or --help
Notes: 		Assumes ./dazy.go exists for helper functions
*/
package main

// Import package dependancies
import (
	"fmt"
	"os"
)

// Define global variables
var (
	progName string = "<REPLACE_NAME>" // Required by ./dazy.go
)

// Desc:
// Usage:
func newFunc() {
	screenOut("INFO", "Started newFunc...")

	// Do stuff
	//	checkStatus(err,
	//		"Error message.",
	//		"Success message.")

	screenOut("INFO", "Finished newFunc.")
}

// Desc: Auto-runs on program execution
// Usage: n/a
func init() {

}

// Desc: Auto-runs after init function
// Usage: n/a
func main() {
	fmt.Println()
	screenOut("INFO", progName+" script started...")

	// Execute functions

	screenOut("INFO", progName+" script ended.")
	fmt.Println()
	os.Exit(0)
}
