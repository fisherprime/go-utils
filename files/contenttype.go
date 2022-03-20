// SPDX-License-Identifier: MIT
package files

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

const (
	// sniffLen defines the length of content to be read from a file to sniff the content type.
	//
	// REF: net/http/sniff.go
	sniffLen = 512
)

// Content type detection errors.
var (
	ErrObtainContentType = errors.New("failed to obtain file content-type")
)

// GetFileContentType retreives the content-type of a file.
//
// Uses `http.DetectContentType` to sniff the content-type from a file.
func GetFileContentType(out *os.File) (contentType string, err error) {
	buffer := make([]byte, sniffLen)

	if _, err = out.Read(buffer); err != nil {
		err = fmt.Errorf("%w: %v", ErrObtainContentType, err)
		return
	}

	// Acquire content-type from the file.
	contentType = http.DetectContentType(buffer)

	return
}
