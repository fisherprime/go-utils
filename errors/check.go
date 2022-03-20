// SPDX-License-Identifier: MIT
package errors

import (
	"log"

	"github.com/joomcode/errorx"
)

// CheckError logs an error message.
//
// Deprecated: This function logs an error message should it be non-`nil` with a user supplied
// message.
// A `true` return indicates a non-`nil` error.
func CheckError(message string, err error) bool {
	if err == nil {
		return false
	}

	if message != "" {
		err = errorx.Decorate(err, message)
	}
	log.Println(err)

	return true
}

// CheckErrorFatal logs an error message then terminates the binary.
//
// Deprecated: This function logs an error should it be non-`nil` with a user supplied error message then exits.
func CheckErrorFatal(message string, err error) {
	if err == nil {
		return
	}

	if message != "" {
		err = errorx.Decorate(err, message)
	}
	log.Fatal(err)
}
