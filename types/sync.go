// SPDX-License-Identifier: MIT
package types

import (
	"sync"
)

type (
	// SafeCounter is a thread-safe counter.
	SafeCounter struct {
		mu    sync.Mutex
		dirty int
	}
)

// Inc increments the counter.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dirty++
}

// Dec decrements the counter.
func (c *SafeCounter) Dec() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dirty--
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.dirty
}
