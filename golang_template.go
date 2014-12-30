/*
  Name:		Replace
  Author:	Daniel Middleton <me@daniel-middleton.com>
  Description:	Replace
  Usage:	Replace
*/

package main

// Import dependancies
import (
	"fmt"
	"os"
	"time"
)

// Global variables
var progName = string("Replace")

// Usage: screenOut("STATUS", "message")
func screenOut(status string, message string) {
	timestamp := time.Now().Format("20060102T150405")

	if status == "" {
		status = "INFO"
	}

	fmt.Printf("[%v][%v][%v]: %v \n", progName, status, timestamp, message)
}

// Usage:
func init() {

}

// Usage:
func main() {
	screenOut("", progName+" script started...")

	// Make a start

	screenOut("", progName+" script ended.")
	os.Exit(0)
}
