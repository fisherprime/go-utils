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
	// Map is a wrapper on the map[string]any type.
	Map map[string]any
)

const (
	// ReadErrFmt defines the format for errors reading a value.
	ReadErrFmt = "failed to read (%s): %w"
	mapErrFmt  = "map %w: %s"
)

// Type errors.
var (
	ErrUndefinedValue = errors.New("undefined value")
	ErrInvalidType    = errors.New("invalid data type")
)

// ReadClaimsErrFmt defines the format for the ErrMissingClaims message.
var ReadClaimsErrFmt = "failed to read claims (%s): %w"

// Add value to Map.
func (a *Map) Add(key string, val any) { (*a)[key] = val }

// Delete from Map.
func (a *Map) Delete(key string) { delete((*a), key) }

// Load from Map.
func (a *Map) Load(key string) (val any, ok bool) {
	val, ok = (*a)[key]
	return
}

// Get value from Map.
//
// Implements the authboss.ClientState interface.
func (a *Map) Get(key string) (out string, ok bool) {
	val, ok := (*a)[key]
	if ok {
		out = fmt.Sprint(val)
	}

	return
}

// LoadString …from Map.
func (a *Map) LoadString(key string, nullable ...bool) (val string, err error) {
	tmp, ok := (*a)[key]
	if !ok || tmp == nil {
		if len(nullable) < 1 || !nullable[0] {
			err = fmt.Errorf(mapErrFmt, ErrUndefinedValue, key)
		}
		return
	}

	if val, ok = tmp.(string); !ok {
		err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
	}

	return
}

// LoadInt from Map.
func (a *Map) LoadInt(key string, nullable ...bool) (val int, err error) {
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

	// Handle frontend requests.
	var id float64
	if id, ok = tmp.(float64); !ok {
		err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
		return
	}
	val = int(id)

	return
}

// LoadUint from Map.
func (a *Map) LoadUint(key string, nullable ...bool) (val uint, err error) {
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

	// Handle frontend requests.
	id, ok := tmp.(float64)
	if !ok {
		err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
		return
	}
	val = uint(id)

	return
}

// LoadBool from Map.
func (a *Map) LoadBool(key string, nullable ...bool) (val bool, err error) {
	tmp, ok := (*a)[key]
	if !ok || tmp == nil {
		if len(nullable) < 1 || !nullable[0] {
			err = fmt.Errorf(mapErrFmt, ErrUndefinedValue, key)
		}
		return
	}

	if val, ok = tmp.(bool); !ok {
		err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
	}

	return
}

// LoadStringSlice obtains a []string from Claims.
func (a *Map) LoadStringSlice(key string, nullable ...bool) (val []string, err error) {
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
		err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
		return
	}

	val = make([]string, len(iSlice))
	for index := range iSlice {
		if val[index], ok = iSlice[index].(string); !ok {
			err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
			return
		}
	}

	return
}

// LoadUintSlice obtains a []uint from Claims.
func (a *Map) LoadUintSlice(key string, nullable ...bool) (val []uint, err error) {
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
		err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
		return
	}

	val = make([]uint, len(iSlice))

	var id float64
	for index := range iSlice {
		if id, ok = iSlice[index].(float64); !ok {
			err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
			return
		}
		val[index] = uint(id)
	}

	return
}

// Merge an Map with the current one.
func (a *Map) Merge(data Map) {
	for k, v := range data {
		(*a)[k] = v
	}
}

// Encode as base64 for [Map].
//
// Using [map[string]any] for compatibility.
func (a *Map) Encode(ctx context.Context) (out []byte, err error) {
	var tmp bytes.Buffer
	if err = gob.NewEncoder(&tmp).Encode((map[string]any)(*a)); err != nil {
		return
	}

	buffer := tmp.Bytes()
	out = make([]byte, base64.StdEncoding.EncodedLen(len(buffer)))

	base64.StdEncoding.Encode(out, buffer)

	return
}

// Decode from base64 for [Map].
//
// Using [map[string]any] for compatibility.
func (a *Map) Decode(ctx context.Context, in []byte) (err error) {
	buffer := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	if _, err = base64.StdEncoding.Decode(buffer, in); err != nil {
		return
	}

	tmp := map[string]any{}
	if err = gob.NewDecoder(bytes.NewReader(buffer)).Decode(&tmp); err != nil {
		return
	}

	*a = Map(tmp)

	return
}
