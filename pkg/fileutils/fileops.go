// Package fileutils provides file utility functions common to myself.
// File modification operations.
package fileutils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gitlab.com/fisherprime/myutils/pkg/errorutils"
)

// PendFileDelete create a waiting job to delete a file.
func PendFileDelete(filePath string) {
	// Check for existing file
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Printf("[!] File does not exist, ", err)

		return
	}

	time.Sleep(120 * time.Second)
	os.Remove(filePath)

	log.Println("[+] Deleted file: " + filePath)
}

// CreateFile creates a file if it doesn't exist.
func CreateFile(filePath string) {
	var err error

	fileDir := filepath.Dir(filePath)

	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, 0755)
		if errorutils.CheckError(fmt.Sprintf("Create directory hierarchy: %s", fileDir), err) {
			return
		}
	}

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if errorutils.CheckError(fmt.Sprintf("Create file: %s", filePath), err) {
			return
		}

		defer file.Close()
	}

	log.Printf("[+] Created file: %s\n", filePath)
}

// WaitUntilFileExists waits for a file to exist before exiting with a nil
// status, otherwise an error should a user specified duration pass before the
// file is available.
// Valid time units: "ns", "us" (or "µs"), "ms", "s", "m", "h"
// Example duration can be "2m".
func WaitUntilFileExists(filePath string, duration string) error {
	var (
		err            error
		parsedDuration time.Duration
	)

	if duration == "" {
		err = errors.New("Empty duration string passed to function")

		return err
	}

	parsedDuration, err = time.ParseDuration(duration)
	if err != nil {
		return err
	}

	stop := time.Now().Add(parsedDuration)

	for {
		_, err := os.Stat(filePath)

		if os.IsNotExist(err) {
			if time.Now().After(stop) {
				return err
			}

			time.Sleep(1 * time.Second)

			continue
		}

		return nil
	}
}
