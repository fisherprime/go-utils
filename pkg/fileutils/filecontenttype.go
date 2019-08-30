// Package fileutils provides file utility functions common to myself.
// File content type functions.
package fileutils

import (
	"net/http"
	"os"

	"gitlab.com/fisherprime/myutils/pkg/errorutils"
)

// GetFileContentType retreives the content-type of a file from it's first 512
// bytes.
func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if errorutils.CheckError("[!] Could not get file content-type, ", err) {
		return "", err
	}

	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
