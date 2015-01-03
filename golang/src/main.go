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
	// Define and parse command line flags
	//wordPtr := flag.String("word", "foo", "a string")
	//numbPtr := flag.Int("numb", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")
	//flag.Parse()

	//fmt.Println("word:", *wordPtr)
	//fmt.Println("numb:", *numbPtr)
	//fmt.Println("fork:", *boolPtr)
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
