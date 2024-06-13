// Package util implements miscellaneous utilities.
//
// SPDX-License-Identifier: MIT
package util

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
