// SPDX-License-Identifier: MIT
package types

import (
	"sync"
)

type (
	// SafeLUT is a thread-safe wrapper for [LUT].
	SafeLUT[K comparable] struct {
		dirty LUT[K]
		m     sync.RWMutex
	}
)

// NewSafeLUT instantiates a `SafeLUT`.
func NewSafeLUT[K comparable]() *SafeLUT[K] {
	return &SafeLUT[K]{
		dirty: make(LUT[K]),
		m:     sync.RWMutex{},
	}
}

// Load from [SafeLUT].
func (l *SafeLUT[K]) Load(key K) (val any, ok bool) {
	l.m.RLock()
	defer l.m.RUnlock()

	val, ok = l.dirty.Load(key)

	return
}

// Store value to Map.
func (l *SafeLUT[K]) Store(key K, val any) {
	l.m.Lock()
	defer l.m.Unlock()

	l.dirty.Store(key, val)
}

// Delete from Map.
func (l *SafeLUT[K]) Delete(key K) {
	l.m.Lock()
	defer l.m.Unlock()

	l.dirty.Delete(key)
}

// Reset [SafeLUT].
func (l *SafeLUT[K]) Reset() {
	l.m.Lock()
	defer l.m.Unlock()

	l.dirty = make(LUT[K])
}

// Merge into [SafeLUT].
func (l *SafeLUT[K]) Merge(lut LUT[K]) {
	l.m.Lock()
	defer l.m.Unlock()

	l.dirty.Merge(lut)
}
