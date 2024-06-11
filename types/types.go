// Package types implements data structures.
//
// SPDX-License-Identifier: MIT
package types

import (
	"errors"
)

// Misc constants.
const (
	UndefinedUint uint = 0
)

// Validation errors.
var (
	ErrInvalidIndex = errors.New("invalid index")
)
