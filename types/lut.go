// SPDX-License-Identifier: MIT
package types

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"log/slog"
	"reflect"
)

type (
	// LUT type definition for a lookup table with a `comparable` key and value of any type.
	LUT[K comparable] map[K]any
)

const (
	// readErrFmt defines the format for errors reading a value.
	readErrFmt = "failed to read (%v): %w"
	mapErrFmt  = "map %w: %v"
)

// Type errors.
var (
	ErrUndefinedValue = errors.New("undefined value")
	ErrInvalidType    = errors.New("invalid data type")
)

// Store value to LUT[K].
func (a *LUT[K]) Store(key K, val any) { (*a)[key] = val }

// Delete from LUT[K].
func (a *LUT[K]) Delete(key K) { delete((*a), key) }

// Load from LUT[K].
func (a *LUT[K]) Load(key K) (val any, ok bool) {
	val, ok = (*a)[key]
	return
}

// LoadString …from LUT[K].
func (a *LUT[K]) LoadString(key K, nullable ...bool) (val string, err error) {
	tmp, ok := (*a)[key]
	if !ok || tmp == nil {
		if len(nullable) < 1 || !nullable[0] {
			err = fmt.Errorf(mapErrFmt, ErrUndefinedValue, key)
		}
		return
	}

	if val, ok = tmp.(string); !ok {
		err = fmt.Errorf(readErrFmt, key, ErrInvalidType)
	}

	return
}

// LoadInt from LUT[K].
func (a *LUT[K]) LoadInt(key K, nullable ...bool) (val int, err error) {
	tmp, ok := (*a)[key]
	if !ok || tmp == nil {
		if len(nullable) < 1 || !nullable[0] {
			err = fmt.Errorf(mapErrFmt, ErrUndefinedValue, key)
		}
		return
	}

	if val, ok = tmp.(int); ok {
		return
	}

	// Should the integer be stored as a float.
	var id float64
	if id, ok = tmp.(float64); !ok {
		err = fmt.Errorf(readErrFmt, key, ErrInvalidType)
		return
	}
	val = int(id)

	return
}

// LoadUint from LUT[K].
func (a *LUT[K]) LoadUint(key K, nullable ...bool) (val uint, err error) {
	tmp, ok := (*a)[key]
	if !ok || tmp == nil {
		if len(nullable) < 1 || !nullable[0] {
			err = fmt.Errorf(mapErrFmt, ErrUndefinedValue, key)
		}
		return
	}

	if val, ok = tmp.(uint); ok {
		return
	}

	// Should the unsigned integer be stored as a float.
	id, ok := tmp.(float64)
	if !ok {
		err = fmt.Errorf(readErrFmt, key, ErrInvalidType)
		return
	}
	val = uint(id)

	return
}

// LoadBool from LUT[K].
func (a *LUT[K]) LoadBool(key K, nullable ...bool) (val bool, err error) {
	tmp, ok := (*a)[key]
	if !ok || tmp == nil {
		if len(nullable) < 1 || !nullable[0] {
			err = fmt.Errorf(mapErrFmt, ErrUndefinedValue, key)
		}
		return
	}

	if val, ok = tmp.(bool); !ok {
		err = fmt.Errorf(readErrFmt, key, ErrInvalidType)
	}

	return
}

// LoadStringSlice obtains a []string from .
func (a *LUT[K]) LoadStringSlice(key K, nullable ...bool) (val []string, err error) {
	tmp, ok := (*a)[key]
	if !ok || tmp == nil {
		if len(nullable) < 1 || !nullable[0] {
			err = fmt.Errorf(mapErrFmt, ErrUndefinedValue, key)
		}
		return
	}

	slog.Debug("LoadStringSlice", "field", key, "type", reflect.TypeOf(tmp), "value", reflect.ValueOf(tmp))

	if val, ok = tmp.([]string); ok {
		return
	}

	iSlice, ok := tmp.([]any)
	if !ok {
		err = fmt.Errorf(readErrFmt, key, ErrInvalidType)
		return
	}

	val = make([]string, len(iSlice))
	for index := range iSlice {
		if val[index], ok = iSlice[index].(string); !ok {
			err = fmt.Errorf(readErrFmt, key, ErrInvalidType)
			return
		}
	}

	return
}

// LoadUintSlice obtains a []uint from .
func (a *LUT[K]) LoadUintSlice(key K, nullable ...bool) (val []uint, err error) {
	tmp, ok := (*a)[key]
	if !ok || tmp == nil {
		if len(nullable) < 1 || !nullable[0] {
			err = fmt.Errorf(mapErrFmt, ErrUndefinedValue, key)
		}
		return
	}

	slog.Debug("LoadUintSlice", "field", key, "type", reflect.TypeOf(tmp), "value", reflect.ValueOf(tmp))

	if val, ok = tmp.([]uint); ok {
		return
	}

	iSlice, ok := tmp.([]any)
	if !ok {
		err = fmt.Errorf(readErrFmt, key, ErrInvalidType)
		return
	}

	val = make([]uint, len(iSlice))

	var id float64
	for index := range iSlice {
		if id, ok = iSlice[index].(float64); !ok {
			err = fmt.Errorf(readErrFmt, key, ErrInvalidType)
			return
		}
		val[index] = uint(id)
	}

	return
}

// Merge an LUT[K] with the current one.
func (a *LUT[K]) Merge(data LUT[K]) {
	for k, v := range data {
		(*a)[k] = v
	}
}

// Encode as base64 for [LUT[K]].
//
// Using [map[string]any] for compatibility.
func (a *LUT[K]) Encode(ctx context.Context) (out []byte, err error) {
	var tmp bytes.Buffer
	if err = gob.NewEncoder(&tmp).Encode((map[K]any)(*a)); err != nil {
		return
	}

	buffer := tmp.Bytes()
	out = make([]byte, base64.StdEncoding.EncodedLen(len(buffer)))

	base64.StdEncoding.Encode(out, buffer)

	return
}

// Decode from base64 for [LUT[K]].
//
// Using [map[string]any] for compatibility.
func (a *LUT[K]) Decode(ctx context.Context, in []byte) (err error) {
	buffer := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	if _, err = base64.StdEncoding.Decode(buffer, in); err != nil {
		return
	}

	tmp := map[K]any{}
	if err = gob.NewDecoder(bytes.NewReader(buffer)).Decode(&tmp); err != nil {
		return
	}

	*a = LUT[K](tmp)

	return
}
