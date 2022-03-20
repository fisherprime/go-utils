// SPDX-License-Identifier: MIT
package types

import (
	"fmt"
	"sort"
	"strings"
)

type (
	// UintSlice is a type wrapper for []uint.
	UintSlice []uint
)

// Locate for UintSlice.
func (sl *UintSlice) Locate(val uint) (resl int) {
	resl = -1

	for index := range *sl {
		if (*sl)[index] == val {
			resl = index
			return
		}
	}

	return
}

// ToStringSlice for UintSlice.
func (sl *UintSlice) ToStringSlice(dst *StringSlice) {
	for index := range *sl {
		(*dst) = append((*dst), fmt.Sprint(((*sl)[index])))
	}
}

// String is the fmt.Stringer implementation for UintSlice.
func (sl *UintSlice) String() (dst string) {
	lenSl := len(*sl)
	if lenSl > 0 {
		buffer := strings.Builder{}
		fmt.Fprintf(&buffer, "[%d", (*sl)[0])
		for index := 1; index < lenSl; index++ {
			fmt.Fprintf(&buffer, ",%d", (*sl)[index])
		}
		buffer.WriteString("]")

		dst = buffer.String()
	}

	return
}

// Sort for UintSlice.
func (sl *UintSlice) Sort() {
	sort.Slice(*sl, func(i, j int) bool { return (*sl)[i] < (*sl)[j] })
}
