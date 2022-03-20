// Package types implements data structures.
//
// SPDX-License-Identifier: MIT
package types

import (
	"errors"

	"github.com/sirupsen/logrus"
)

var (
	// fLogger is a logrus.FieldLogger used for debug purposes.
	fLogger logrus.FieldLogger = logrus.NewEntry(logrus.New())
)

// Validation errors.
var (
	ErrInvalidIndex = errors.New("invalid index")
)

// SetLogger sets the debugging logrus.FieldLogger.
func SetLogger(l *logrus.Entry) { fLogger = l }
