// SPDX-License-Identifier: MIT
package types

import "encoding/json"

// AppendToAny using the lazy approach.
//
// Expects a pointer to a type for the destination.
func AppendToAny(src, dst any) (err error) {
	buffer, err := json.Marshal(src)
	if err != nil {
		return
	}
	err = json.Unmarshal(buffer, dst)

	return
}
