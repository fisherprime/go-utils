// Package errorutils provides error utility functions common to myself
package errorutils

import (
	"log"

	"github.com/joomcode/errorx"
)

// CheckError checks the error variable; if set, prints out the error message
// appended to a user specified-string --or null-- then returns true, else
// returns false.
// The user-specified string says what operation was occurring: Unmarshall
// MyType struct, Parse thisValue, ...
func CheckError(message string, err error) bool {
	if err != nil {
		if message != "" {
			log.Println(errorx.Decorate(err, message))
			return (err != nil)
		}
		log.Println(err)
	}

	return (err != nil)
}

// CheckErrorFatal checks the error variable, if set prints out the error
// message appended to a user chosen string --or null-- then exits the
// application.
// The user-specified string says what operation was occurring: Unmarshall
// MyType struct, Parse thisValue, ...
func CheckErrorFatal(message string, err error) {
	if err != nil {
		if message != "" {
			log.Fatal(errorx.Decorate(err, message))
		}
		log.Fatal(err)
	}
}
