// Package errorutils provides error utility functions common to myself
package errorutils

import (
	log "github.com/sirupsen/logrus"
)

// CheckError checks for a valid error, returing true for the case & false otherwise.
//
// If an error is valid, this method prints out the error message appended to a user chosen string.
func CheckError(message string, err error) (resl bool) {
	if err == nil {
		return
	}

	resl = true
	if message != "" {
		log.WithError(err).Error(message)
		return
	}

	log.Error(err)

	return
}

// CheckErrorFatal checks for a valid error then terminates the application.
//
// If an error is valid, this method prints out the error message appended to a user chosen string.
func CheckErrorFatal(message string, err error) {
	if err == nil {
		return
	}

	if message != "" {
		log.WithError(err).Fatal(message)
		return
	}

	log.Fatal(err)
}
