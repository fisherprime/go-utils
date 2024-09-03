// SPDX-License-Identifier: MIT
package types

import "encoding/json"

// AppendToAny using json.Marshal & json.Unmarshal (lazy approach).
//
// Expects a pointer as the destination.
func AppendToAny(src, dst any) (err error) {
	buffer, err := json.Marshal(src)
	if err != nil {
		return
	}

	return json.Unmarshal(buffer, dst)
}
