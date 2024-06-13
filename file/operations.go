// SPDX-License-Identifier: MIT
package file

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	defaultFilePerm = 0o600
	defaultDirPerm  = 0o700
)

// File operation errors.
var (
	ErrFailedToOpenFile   = errors.New("failed to open file")
	ErrCreateDirHierarchy = errors.New("failed to create directory hierarchy")
)

// WaitUntilFileExists polls until context expiry for a file to exist.
func WaitUntilFileExists(ctx context.Context, filePath string) (err error) {
	if filePath, err = filepath.Abs(filePath); err != nil {
		return
	}

	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		var t *time.Timer
		for {
			if _, err = os.Stat(filePath); err != nil {
				if !os.IsNotExist(err) {
					return
				}

				t = time.NewTimer(1 * time.Second)
				<-t.C

				continue
			}

			break
		}
	}

	return
}

// WriteFile overwrites the contents of a file with the supplied data.
//
// Creates a file if it did not exist together with its directory hierarchy.
func WriteFile(ctx context.Context, filePath string, data []byte) (err error) {
	if filePath, err = filepath.Abs(filePath); err != nil {
		return
	}

	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		if err = os.MkdirAll(path.Dir(filePath), defaultDirPerm); err != nil {
			return fmt.Errorf("%w: %v", ErrCreateDirHierarchy, err)
		}

		if err = os.WriteFile(filePath, data, defaultFilePerm); err != nil {
			err = fmt.Errorf("%w (%s)", err, filePath)
		}
	}

	return
}

// ReadFile reads the contents of a file.
func ReadFile(ctx context.Context, filePath string) (b []byte, err error) {
	if filePath, err = filepath.Abs(filePath); err != nil {
		return
	}

	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		if b, err = os.ReadFile(filePath); err != nil {
			err = fmt.Errorf("%w (%s)", err, filePath)
		}
	}

	return
}
