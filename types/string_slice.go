package types

import (
	"sort"
	"strconv"
)

type (
	// ByteSlice is a type wrapper for []byte.
	ByteSlice []byte

	// StringSlice is a type wrapper for []string.
	StringSlice []string
)

// Sort for StringSlice.
func (sl *StringSlice) Sort() {
	sort.Strings(*sl)
	// sort.Slice(*sl, func(i, j int) bool {return ()})
}

// Locate for StringSlice.
func (sl *StringSlice) Locate(val string) (resl int) {
	resl = -1

	for index := range *sl {
		if (*sl)[index] == val {
			resl = index
			return
		}
	}

	return
}

// UniqueAppend to StringSlice.
func (sl *StringSlice) UniqueAppend(values ...string) {
	if len(values) < 1 {
		return
	}

	for index := range values {
		newValue := values[index]
		if sl.Locate(newValue) > -1 {
			continue
		}

		*sl = append(*sl, newValue)
	}
}

// UniquePrepend to StringSlice.
func (sl *StringSlice) UniquePrepend(values ...string) {
	if len(values) < 1 {
		return
	}

	for index := range values {
		newValue := values[index]
		if sl.Locate(newValue) > -1 {
			continue
		}

		*sl = append(StringSlice{newValue}, *sl...)
	}
}

// ToByteSlice converts a StringSlice to a [][]byte.
func (sl *StringSlice) ToByteSlice(output *[][]byte) {
	for index := range *sl {
		*output = append(*output, ByteSlice((*sl)[index]))
	}
}

// ToUintSlice for StringSlice.
func (sl *StringSlice) ToUintSlice(dst *UintSlice) (err error) {
	var val uint64
	for index := range *sl {
		if val, err = strconv.ParseUint((*sl)[index], 10, 64); err != nil {
			return
		}
		(*dst) = append((*dst), uint(val))
	}

	return
}

// CutValues from StringSlice.
func (sl *StringSlice) CutValues(values ...string) {
	for index := range values {
		if loc := sl.Locate(values[index]); loc > -1 {
			sl.Cut(loc)
		}
	}
}

// Cut index from StringSlice.
func (sl *StringSlice) Cut(index int) {
	upper := index + 1

	// index == 0
	if index < 1 {
		// NOTE: For clarity, using the length without a check will yield an empty slice.
		if upper >= len(*sl) {
			*sl = StringSlice{}
			return
		}

		*sl = (*sl)[1:]
		return
	}

	*sl = append((*sl)[:index], (*sl)[upper:]...)
}
