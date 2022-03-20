// SPDX-License-Identifier: MIT
package files

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// File operation errors.
var (
	ErrFailedToOpenFile   = errors.New("failed to open file")
	ErrCreateDirHierarchy = errors.New("failed to create directory hierarchy")
	ErrCreateFile         = errors.New("failed to create file")
)

// PendFileDelete creates a waiting job to delete a file.
//
// Deprecated: Takes a filepath & `time.Duration`: `("my/cool/path", 2 * time.Second)`.
func PendFileDelete(_ context.Context, filePath string, duration time.Duration) (err error) {
	// Check for existing file
	if _, err = os.Stat(filePath); err != nil {
		/* if os.IsNotExist(err) {
		 *     log.Printf("file does not exist, %v", err)
		 *     return
		 * } */

		return
	}

	t := time.NewTimer(duration)
	<-t.C
	if err = os.Remove(filePath); err != nil {
		return
	}

	// log.Println("deleted file: " + filePath)

	return
}

// CreateFile creates a file if it doesn't exist.
//
// Deprecated: .
func CreateFile(ctx context.Context, filePath string) (err error) {
	if filePath, err = filepath.Abs(filePath); err != nil {
		return
	}

	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:

		fileDir := filepath.Dir(filePath)

		if _, err = os.Stat(fileDir); err != nil {
			if !os.IsNotExist(err) {
				return
			}

			if err = os.MkdirAll(fileDir, 0755); err != nil {
				err = fmt.Errorf("%w: %v", ErrCreateDirHierarchy, err)
				return
			}
		}

		if _, err = os.Stat(filePath); err != nil {
			if !os.IsNotExist(err) {
				return
			}

			var file *os.File
			if file, err = os.Create(filePath); err != nil {
				err = fmt.Errorf("%w: %v", ErrCreateFile, err)
				return
			}
			err = file.Close()
		}

		// log.Printf("created file: %s\n", filePath)
	}

=======
	fileDir := filepath.Dir(filePath)

	if _, err = os.Stat(fileDir); err != nil {
		if !os.IsNotExist(err) {
			return
		}

		if err = os.MkdirAll(fileDir, 0755); err != nil {
			err = fmt.Errorf("%w: %v", ErrCreateDirHierarchy, err)
			return
		}
	}

	if _, err = os.Stat(filePath); err != nil {
		if !os.IsNotExist(err) {
			return
		}

		var file *os.File
		if file, err = os.Create(filePath); err != nil {
			err = fmt.Errorf("%w: %v", ErrCreateFile, err)
			return
		}
		err = file.Close()
	}

	// log.Printf("created file: %s\n", filePath)

>>>>>>> 19e9fc1ce4f43ce1574a95a41774f257ab3c35bb
	return
}

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

// OverwriteFile overwrites the contents of a file with the supplied data.
func OverwriteFile(ctx context.Context, filePath string, data []byte) (err error) {
	if filePath, err = filepath.Abs(filePath); err != nil {
		return
	}

	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		var file *os.File

		// nolint: gosec // "G304" is sorted out by `filepath.Abs`.
		if file, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0600); err != nil {
			err = fmt.Errorf("%w %s: %v", ErrFailedToOpenFile, filePath, err)
			return
		}

		_, err = file.Write(data)
		_ = file.Close()
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
		var file *os.File

		// nolint: gosec // "G304" is sorted out by `filepath.Abs`.
		if file, err = os.OpenFile(filePath, os.O_RDONLY, 0600); err != nil {
			err = fmt.Errorf("%w %s: %v", ErrFailedToOpenFile, filePath, err)
			return
		}

		b, err = io.ReadAll(file)
		_ = file.Close()
	}

	return
}
