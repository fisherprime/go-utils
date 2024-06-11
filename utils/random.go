// SPDX-License-Identifier: MIT
package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (

	// randStringAlphabet contains the alphabet used in generating the random string.
	randStringAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"

	// secretLen is the default length for the system's secrets.
	secretLen = 32
)

// nolint: gosec // G404 is not critical for this implementation.
var rSeed = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateSecret generates a secret (random string).
func GenerateSecret() string { return GenerateRandString(secretLen) }

// GenerateRandString generates a random string.
func GenerateRandString(sLen int) string {
	strBuilder := strings.Builder{}
	strBuilder.Grow(sLen)

	lenAlphabet := len(randStringAlphabet)

	for index := sLen - 1; index >= 0; index-- {
		strBuilder.WriteByte(randStringAlphabet[rand.Intn(lenAlphabet)]) // nolint: gosec
	}

	return strBuilder.String()
}
