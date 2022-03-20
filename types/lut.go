// SPDX-License-Identifier: MIT
package types

import (
	"sync"
)

// TODO: Look into using interface{}s.
type (
	// UStringLUT is a map[uint]string lookup table.
	UStringLUT map[uint]string
	// SUintLUT is a map[string]uint lookup table.
	SUintLUT map[string]uint

	// SStringLUT is a map[string]string lookup table.
	SStringLUT map[string]string

	// SafeSStringLUT is a thread-safe SStringLUT.
	SafeSStringLUT struct {
		mu    sync.RWMutex
		dirty SStringLUT
	}

	// SafeSUintLUT is a thread-safe SUintLUT.
	SafeSUintLUT struct {
		mu    sync.RWMutex
		dirty SUintLUT
	}

	// SafeUStringLUT is a thread-safe UStringLUT.
	SafeUStringLUT struct {
		mu    sync.RWMutex
		dirty UStringLUT
	}
)

// Put to SafeValLUT.
func (l *SafeUStringLUT) Put(k uint, v string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.dirty[k] = v
}

// Put to SafeIDLUT.
func (l *SafeSUintLUT) Put(k string, v uint) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.dirty[k] = v
}

// Put to SafeValLUT.
func (l *SafeSStringLUT) Put(k, v string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.dirty[k] = v
}

// Get from SafeValLUT.
func (l *SafeUStringLUT) Get(k uint) (v string, ok bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	v, ok = l.dirty[k]

	return
}

// Get from SafeIDLUT.
func (l *SafeSUintLUT) Get(k string) (v uint, ok bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	v, ok = l.dirty[k]

	return
}

// Get from SafeStringLUT.
func (l *SafeSStringLUT) Get(k string) (v string, ok bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	v, ok = l.dirty[k]

	return
}

// Reset SafeIDLUT.
func (l *SafeSUintLUT) Reset() {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.dirty = make(SUintLUT)
}

// Reset SafeValLUT.
func (l *SafeUStringLUT) Reset() {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.dirty = make(UStringLUT)
}

// Reset SafeStringLUT.
func (l *SafeSStringLUT) Reset() {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.dirty = make(SStringLUT)
}

// Set SafeIDLUT.
func (l *SafeSUintLUT) Set(lut SUintLUT) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.dirty = lut
}

// Set SafeValLUT.
func (l *SafeUStringLUT) Set(lut UStringLUT) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.dirty = lut
}

// Set SafeStringLUT.
func (l *SafeSStringLUT) Set(lut SStringLUT) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.dirty = lut
}
