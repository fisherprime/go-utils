// SPDX-License-Identifier: MIT
package types

import (
	"fmt"

	"github.com/lib/pq"
)

// StoreIDSlice is a type wrapper for pq.Int64Array.
type StoreIDSlice pq.Int64Array

// ToStringSlice for StoreIDList.
func (sl *StoreIDSlice) ToStringSlice(dst *StringSlice) {
	for index := range *sl {
		(*dst) = append((*dst), fmt.Sprint(((*sl)[index])))
	}
}

// Locate checks for the existence of a store-type ID.
func (sl *StoreIDSlice) Locate(n uint) (resl int) {
	role := int64(n)
	for index := range *sl {
		if (*sl)[index] == role {
			resl = index
			return
		}
	}

	return
}
