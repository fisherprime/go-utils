// Package fileutils provides file utility functions common to myself.
// File modification operations.
package fileutils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gitlab.com/fisherprime/myutils/pkg/errorutils"
)

// PendFileDelete create a waiting job to delete a file.
// Takes a filepath & time.Duration:  ("my/cool/path", 2 * time.Second)
func PendFileDelete(filePath string, duration time.Duration) error {
	// Check for existing file
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Printf("[!] File does not exist, ", err)

		return err
	}

	time.Sleep(duration)
	os.Remove(filePath)

	log.Println("[+] Deleted file: " + filePath)
	return nil
}

// CreateFile creates a file if it doesn't exist.
func CreateFile(filePath string) error {
	var err error

	fileDir := filepath.Dir(filePath)

	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, 0755)
		if errorutils.CheckError(fmt.Sprintf("Create directory hierarchy: %s", fileDir), err) {
			return err
		}
	}

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if errorutils.CheckError(fmt.Sprintf("Create file: %s", filePath), err) {
			return err
		}

		defer file.Close()
	}

	log.Printf("[+] Created file: %s\n", filePath)
	return nil
}

// WaitUntilFileExists waits for a file to exist before exiting with a nil
// status, otherwise an error should a user specified duration pass before the
// file is available.
// Example duration: 2 * time.Second
func WaitUntilFileExists(filePath string, duration time.Duration) error {
	// Valid time units: "ns", "us" (or "µs"), "ms", "s", "m", "h"
	stop := time.Now().Add(duration)

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
