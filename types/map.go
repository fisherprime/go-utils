// SPDX-License-Identifier: MIT
package types

import (
	"errors"
	"fmt"
	"reflect"
)

type (
	// Map wraps the map[string]interface{} type.
	//
	// This replicates functionality in sync.Map without synchronization.
	// Not safe for concurrent use.
	Map map[string]interface{}
)

const (
	// ReadErrFmt defines the format string for failure to read data.
	ReadErrFmt = "failed to read (%s): %w"
)

// Type related errors.
var (
	ErrInvalidType = errors.New("invalid data type")
)

// Store value to Map.
func (a *Map) Store(key string, value interface{}) { (*a)[key] = value }

// Delete from Map.
func (a *Map) Delete(key string) { delete((*a), key) }

// Load from Map.
func (a *Map) Load(key string) (out interface{}, ok bool) {
	if out = (*a)[key]; out != nil {
		ok = true
	}

	return
}

// LoadAndDelete from Map.
func (a *Map) LoadAndDelete(key string) (value interface{}, loaded bool) {
	if value = (*a)[key]; value != nil {
		loaded = true
		a.Delete(key)
	}

	return
}

// LoadOrStore key,value pair to Map.
func (a *Map) LoadOrStore(key string, value interface{}) (actual interface{}, loaded bool) {
	if actual, loaded = a.Load(key); loaded {
		return
	}
	a.Store(key, value)

	return
}

// LoadString from the Map.
func (a *Map) LoadString(key string) (strVal string, err error) {
	if val, ok := (*a)[key]; ok {
		if strVal, ok = val.(string); !ok {
			err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
		}
	}

	return
}

// LoadInt from the Map.
func (a *Map) LoadInt(key string) (value int, err error) {
	if val, ok := (*a)[key]; ok {
		var id float64
		if id, ok = val.(float64); !ok {
			err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
			return
		}
		value = int(id)
	}

	return
}

// LoadlUint from the Map.
func (a *Map) LoadlUint(key string) (value uint, ok bool, err error) {
	var val interface{}
	if val, ok = (*a)[key]; ok {
		// Handle frontend requests.
		var id float64
		if id, ok = val.(float64); !ok {
			// Handle authorization claims.
			if value, ok = val.(uint); !ok {
				err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
				return
			}
			return
		}
		value = uint(id)
	}

	return
}

// LoadBool from the Map.
func (a *Map) LoadBool(key string) (value bool, err error) {
	if val, ok := (*a)[key]; ok {
		if value, ok = val.(bool); !ok {
			err = fmt.Errorf(ReadErrFmt, key, ErrInvalidType)
		}
	}

	return
}

// LoadStringSlice obtains an []interface{} from the Map.
func (a *Map) LoadStringSlice(fieldName string) (value StringSlice, err error) {
	val, ok := (*a)[fieldName]
	if !ok || val == nil {
		return
	}

	fLogger.Debugf("[]string field: %s, type: %v, value: %v", fieldName, reflect.TypeOf(val), reflect.ValueOf(val))

	var iSlice []interface{}
	if iSlice, ok = val.([]interface{}); !ok {
		err = fmt.Errorf(ReadErrFmt, fieldName, ErrInvalidType)
		return
	}

	value = make(StringSlice, len(iSlice))
	for index := range iSlice {
		if value[index], ok = iSlice[index].(string); !ok {
			err = fmt.Errorf(ReadErrFmt, fieldName, ErrInvalidType)
			return
		}
	}

	return
}

// LoadUintSlice obtains a []uint from the Map.
func (a *Map) LoadUintSlice(fieldName string) (value UintSlice, err error) {
	val, ok := (*a)[fieldName]
	if !ok || val == nil {
		return
	}

	fLogger.Debugf("[]uint field: %s, type: %v, value: %v", fieldName, reflect.TypeOf(val), reflect.ValueOf(val))

	var iSlice []interface{}
	if iSlice, ok = val.([]interface{}); !ok {
		err = fmt.Errorf(ReadErrFmt, fieldName, ErrInvalidType)
		return
	}

	value = make(UintSlice, len(iSlice))

	var id float64
	for index := range iSlice {
		if id, ok = iSlice[index].(float64); !ok {
			err = fmt.Errorf(ReadErrFmt, fieldName, ErrInvalidType)
			return
		}
		value[index] = uint(id)
	}

	return
}

// Merge a Map with the current one.
func (a *Map) Merge(data Map) {
	for k, v := range data {
		(*a)[k] = v
	}
}
