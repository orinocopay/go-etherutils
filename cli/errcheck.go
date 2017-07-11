package cli

import (
	"fmt"
	"os"
)

// ErrCheck checks for an error and quits if it is present
func ErrCheck(err error, quiet bool, msg string) {
	if err != nil {
		if !quiet {
			fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err.Error())
		}
		os.Exit(1)
	}
}

// Assert checks a condition and quits if it is false
func Assert(condition bool, quiet bool, msg string) {
	if !condition {
		Err(quiet, msg)
	}
}

// Err prints an erro rand quits
func Err(quiet bool, msg string) {
	if !quiet {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
	os.Exit(1)
}
