// Package myUtils provides utility fuctions common to myself
package myUtils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// GetFileContentType retreives the content-type of a file from it's first 512
// bytes.
func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if CheckError("[!] Could not get file content-type, ", err) {
		return "", err
	}

	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

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

// CheckError checks the error variable; if set, prints out the error message
// appended to a user specified string --or null-- then returns true, else
// returns false.
func CheckError(message string, err error) bool {
	if err != nil {
		log.Println(fmt.Sprintf("%s,", message), err)
	}

	return (err != nil)
}

// CheckErrorFatal checks the error variable, if set prints out the error
// message appended to a user chosen string --or null-- then exits the
// application.
func CheckErrorFatal(message string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s,", message), err)
	}
}

// CreateFile creates a file if it doesn't exist.
func CreateFile(filePath string) {
	var err error

	fileDir := filepath.Dir(filePath)

	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(fileDir, 0644)
		if CheckError(fmt.Sprintf("[!] Error occurred while creating directory hierarchy: %s", fileDir), err) {
			return
		}
	}

	_, err = os.Stat(filePath)

	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if CheckError(fmt.Sprintf("[!] Error occurred while creating file: %s", filePath), err) {
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
