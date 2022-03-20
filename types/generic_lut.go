// SPDX-License-Identifier: MIT
package types

import "sync"

type (
	// GenericLut implements a lookup table using a comparable generic type for the key.
	GenericLut[K comparable] struct {
		Values map[K]interface{}
	}

	// ComparableLut implements a lookup table using a comparable generic type for the key & value.
	ComparableLut[K comparable, V comaprable] struct {
		Values map[K]V
	}

	SafeGenericLut[K comparable, V comparable] struct {
		mu    sync.RWMutex
		dirty ComparableLut[K, V]
	}
)

// NewGenericLut instantiates a GenericLut.
func NewGenericLut[T comparable]() GenericLut[T] { return GenericLut{Values: make(map[T]interface{})} }

// NewComparableLut instantiates a ComparableLut.
func NewComparableLut[T comparable]() GenericLut[K, V] { return ComparableLut{Values: make(map[K]V)} }

func NewSafeComparableLut[T comparable]() GenericLut[K, V] {
	return ComparableLut{Values: make(map[K]V)}
}

// Put to SafeValLUT.
func (l *SafeSStringLUT) Put(k, v comparable) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.dirty[k] = v
}

// Get from SafeComparableLut.
func (l *SafeComparableLut) Get(k comparable) (v comparable, ok bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	v, ok = l.dirty[k]

	return
}

// Reset SafeComparableLut.
func (l *SafeComparableLut) Reset() {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var kType, vType reflect.Type
	for k,v := l.dirty {
		kType, vType = reflect.TypeOf(k), reflect.TypeOf(v)
		break
	}

	l.dirty = make(SafeComparableLut[kType,vType])
}

// Set SafeComparableLut.
func (l *SafeComparableLut) Set(lut ComparableLut) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	l.dirty = lut
}
