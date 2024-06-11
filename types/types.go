// Package types implements data structures.
//
// SPDX-License-Identifier: MIT
package types

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// Misc constants.
const (
	UndefinedUint uint = 0
)

// fLogger is a logrus.FieldLogger used for debug purposes.
var fLogger logrus.FieldLogger = logrus.NewEntry(logrus.New())

// Validation errors.
var (
	ErrInvalidIndex = errors.New("invalid index")
)

// SetLogger sets the debugging logrus.FieldLogger.
func SetLogger(l *logrus.Entry) { fLogger = l }
