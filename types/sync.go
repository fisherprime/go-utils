// SPDX-License-Identifier: MIT
package types

import (
	"sync"
)

type (
	// SafeCounter is a thread-safe counter.
	SafeCounter struct {
		val int
		m   sync.Mutex
	}
)

// Inc increments the counter.
func (c *SafeCounter) Inc() {
	c.m.Lock()
	defer c.m.Unlock()
	c.val++
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.val
}
