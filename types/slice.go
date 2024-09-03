package types

import (
	"cmp"
	"fmt"
	"sort"
	"strings"
)

// Slice is a type wrapper for []T where T satisfies the cmp.Ordered constraint.
//
// swagger:model [slice]
type Slice[T cmp.Ordered] []T

// Locate for [Slice].
func (sl *Slice[T]) Locate(val T) (loc int) {
	loc = -1

	for index := range *sl {
		if (*sl)[index] == val {
			loc = index
			return
		}
	}

	return
}

// String is the [fmt.Stringer] implementation for [Slice].
func (sl *Slice[T]) String() (dst string) {
	lenSl := len(*sl)
	if lenSl < 1 {
		return
	}

	buffer := strings.Builder{}
	fmt.Fprint(&buffer, "[", (*sl)[0])
	for index := 1; index < lenSl; index++ {
		fmt.Fprint(&buffer, ",", (*sl)[index])
	}
	buffer.WriteString("]")

	dst = buffer.String()

	return
}

// ToCache for [Slice].
func (sl *Slice[T]) ToCache() (cache map[T]struct{}) {
	cache = make(map[T]struct{})
	for _, val := range *sl {
		cache[val] = struct{}{}
	}

	return
}

// Append for [Slice].
func (sl *Slice[T]) Append(values ...T) {
	if len(values) < 1 {
		return
	}

	for index := range values {
		*sl = append(*sl, values[index])
	}
}

// Prepend to [Slice].
func (sl *Slice[T]) Prepend(values ...T) {
	if len(values) < 1 {
		return
	}

	for index := range values {
		*sl = append(Slice[T]{values[index]}, *sl...)
	}
}

// UniqueAppend for [Slice].
func (sl *Slice[T]) UniqueAppend(values ...T) (appended []T) {
	if len(values) < 1 {
		return
	}

	cache := sl.ToCache()

	appended = Slice[T]{}
	for index := range values {
		newVal := values[index]
		if _, ok := cache[newVal]; ok {
			continue
		}

		*sl = append(*sl, newVal)
		appended = append(appended, newVal)
		cache[newVal] = struct{}{}
	}

	return
}

// UniquePrepend to [Slice].
func (sl *Slice[T]) UniquePrepend(values ...T) (prepended []T) {
	if len(values) < 1 {
		return
	}

	cache := sl.ToCache()

	prepended = Slice[T]{}
	for index := range values {
		newVal := values[index]
		if _, ok := cache[newVal]; ok {
			continue
		}

		*sl = append(Slice[T]{newVal}, *sl...)
		prepended = append(prepended, newVal)
		cache[newVal] = struct{}{}
	}

	return
}

// Pop from [Slice].
func (sl *Slice[T]) Pop(index int) {
	base := index
	upper := index + 1
	if base < 1 {
		// NOTE: Availed for clarity, using the length before the colon will yield an empty slice.
		if upper >= len(*sl) {
			*sl = Slice[T]{}
			return
		}

		*sl = (*sl)[1:]
		return
	}

	*sl = append((*sl)[:base], (*sl)[upper:]...)
}

// PopValues from [Slice].
func (sl *Slice[T]) PopValues(values ...T) {
	lenValues := len(values)
	switch {
	case lenValues < 1, len(*sl) < 1:
		return
	case lenValues == 1:
		if loc := sl.Locate(values[0]); loc > -1 {
			sl.Pop(loc)
		}
		return
	default:
	}

	cache := sl.ToCache()
	for _, value := range values {
		delete(cache, value)
	}

	lenCache := len(cache)
	newSl := make(Slice[T], lenCache)

	if lenCache > 0 {
		for newIndex, index := 0, 0; index < lenCache; index++ {
			value := (*sl)[index]
			if _, ok := cache[value]; !ok {
				continue
			}

			newSl[newIndex] = value
			newIndex++
		}
	}

	*sl = newSl
}

// Sort for [Slice].
func (sl *Slice[T]) Sort() {
	sort.Slice(*sl, func(i, j int) bool { return (*sl)[i] < (*sl)[j] })
}
