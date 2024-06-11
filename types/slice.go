package types

import (
	"cmp"
	"fmt"
	"sort"
	"strings"
)

// Slice is a wrapper on [slice] for [cmp.Ordered] values.
//
// swagger:model [slice]
type Slice[V cmp.Ordered] []V

// Locate for [Slice].
func (sl *Slice[V]) Locate(val V) (loc int) {
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
func (sl *Slice[V]) String() (dst string) {
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
func (sl *Slice[V]) ToCache() (cache map[V]struct{}) {
	cache = make(map[V]struct{})
	for _, val := range *sl {
		cache[val] = struct{}{}
	}

	return
}

// Append for [Slice].
func (sl *Slice[V]) Append(values ...V) {
	if len(values) < 1 {
		return
	}

	for index := range values {
		*sl = append(*sl, values[index])
	}
}

// Prepend to [Slice].
func (sl *Slice[V]) Prepend(values ...V) {
	if len(values) < 1 {
		return
	}

	for index := range values {
		*sl = append(Slice[V]{values[index]}, *sl...)
	}
}

// UniqueAppend for [Slice].
func (sl *Slice[V]) UniqueAppend(values ...V) (appended []V) {
	if len(values) < 1 {
		return
	}

	cache := sl.ToCache()

	appended = Slice[V]{}
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
func (sl *Slice[V]) UniquePrepend(values ...V) (prepended []V) {
	if len(values) < 1 {
		return
	}

	cache := sl.ToCache()

	prepended = Slice[V]{}
	for index := range values {
		newVal := values[index]
		if _, ok := cache[newVal]; ok {
			continue
		}

		*sl = append(Slice[V]{newVal}, *sl...)
		prepended = append(prepended, newVal)
		cache[newVal] = struct{}{}
	}

	return
}

// Pop from [Slice].
func (sl *Slice[V]) Pop(index int) {
	base := index
	upper := index + 1
	if base < 1 {
		// NOTE: Availed for clarity, using the length before the colon will yield an empty slice.
		if upper >= len(*sl) {
			*sl = Slice[V]{}
			return
		}

		*sl = (*sl)[1:]
		return
	}

	*sl = append((*sl)[:base], (*sl)[upper:]...)
}

// PopValues from [Slice].
func (sl *Slice[V]) PopValues(values ...V) {
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
	newSl := make(Slice[V], lenCache)

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
func (sl *Slice[V]) Sort() {
	sort.Slice(*sl, func(i, j int) bool { return (*sl)[i] < (*sl)[j] })
}
