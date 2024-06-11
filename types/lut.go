// SPDX-License-Identifier: MIT
package types

import (
	"sync"
)

type (
	// LUT …
	LUT[K comparable, V any] map[K]V

	// SafeLUT is a thread-safe wrapper for [LUT].
	SafeLUT[K comparable, V any] struct {
		dirty LUT[K, V]
		m     sync.RWMutex
	}
)

// Put into [SafeLUT].
func (l *SafeLUT[K, V]) Put(lut LUT[K, V]) {
	l.m.Lock()
	defer l.m.Unlock()

	l.dirty = lut
}

// Get from [SafeLUT].
func (l *SafeLUT[K, V]) Get(id K) (val V, ok bool) {
	l.m.RLock()
	defer l.m.RUnlock()

	val, ok = l.dirty[id]

	return
}

// Reset [SafeLUT].
func (l *SafeLUT[K, V]) Reset() {
	l.m.Lock()
	defer l.m.Unlock()

	l.dirty = make(LUT[K, V])
}
