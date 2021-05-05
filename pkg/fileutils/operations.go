// Package fileutils provides file utility functions common to myself.
// File modification operations.
package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

// PendFileDelete create a waiting job to delete a file.
// Takes a filepath & time.Duration:  ("my/cool/path", 2 * time.Second)
func PendFileDelete(filePath string, duration time.Duration) (err error) {
	// Check for existing file
	if _, err = os.Stat(filePath); err != nil {
		return
	}

	time.Sleep(duration)
	if err = os.Remove(filePath); err != nil {
		return
	}

	log.Info("Deleted file: " + filePath)

	return
}

// CreateFile creates a file if it doesn't exist.
func CreateFile(filePath string) (err error) {
	fileDir := filepath.Dir(filePath)

	if _, err = os.Stat(fileDir); err != nil {
		if !os.IsNotExist(err) {
			return
		}

		if err = os.MkdirAll(fileDir, 0755); err != nil {
			err = fmt.Errorf("create directory hierarchy: %s, %w", fileDir, err)
			return
		}
	}

	if _, err = os.Stat(filePath); err != nil {
		if !os.IsNotExist(err) {
			return
		}

		var file *os.File
		if file, err = os.Create(filePath); err != nil {
			err = fmt.Errorf("create file: %s, %w", filePath, err)
			return
		}

		if err = file.Close(); err != nil {
			return
		}
	}

	log.Info("Created file: %s\n", filePath)

	return nil
}

// WaitUntilFileExists waits for a file to exist.
//
// An error is returned should the user specified timout expire before the file is available.
func WaitUntilFileExists(filePath string, timeout time.Duration) (err error) {
	// Valid time units: "ns", "us" (or "µs"), "ms", "s", "m", "h"
	stop := time.Now().Add(timeout)

	for {
		if _, err = os.Stat(filePath); err != nil {
			if !os.IsNotExist(err) {
				return
			}
		}

		if time.Now().After(stop) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}
