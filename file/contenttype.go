// SPDX-License-Identifier: MIT
package file

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

const (
	// mimeSniffLength is the number of bytes necessary to determine a file's content type by
	// `http.DetectContentType`.
	//
	// REF: net/http/sniff.go
	mimeSniffLength = 512
)

// Content type detection errors.
var (
	ErrObtainContentType = errors.New("failed to obtain file content-type")
)

// GetFileContentType retreives the content-type of a file.
//
// Uses `http.DetectContentType` to sniff the content-type from a file.
func GetFileContentType(file *os.File) (contentType string, err error) {
	buffer := make([]byte, mimeSniffLength)

	if _, err = file.Read(buffer); err != nil {
		err = fmt.Errorf("%w: %w", ErrObtainContentType, err)

		return
	}

	// Acquire content-type from the file.
	contentType = http.DetectContentType(buffer)

	return
}
