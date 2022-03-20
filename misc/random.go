// SPDX-License-Identifier: MIT
package misc

import (
	"math/rand"
	"time"
)

const (
	// RandomStringAlphabet contains the alphabet used in generating random strings.
	RandomStringAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"

	// secretLen is the default length for the system's secrets.
	secretLen = 32
)

var (
	// nolint: gosec // G404 is not critical for this implementation.
	rSeed = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// GenSecret generates a secret (random string).
func GenSecret() string { return RandomStringFromAlphabet(secretLen, RandomStringAlphabet) }

// RandomStringFromAlphabet generates a random string.
func RandomStringFromAlphabet(length int, alphabet string) string {
	buffer := make([]byte, length)

	lenAlphabet := len(alphabet)
	for index := range buffer {
		buffer[index] = alphabet[rSeed.Intn(lenAlphabet)]
	}

	return string(buffer)
}
