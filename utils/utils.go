// Package utils implements miscellaneous utilities.
//
// SPDX-License-Identifier: MIT
package utils

import (
	"errors"
)

const (
	// ErrChanSize defines the size for an error channel.
	ErrChanSize = 5
)

// Synchronization errors.
var (
	ErrInvalidGoroutineCount = errors.New("invalid goroutine count")
)
