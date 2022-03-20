// SPDX-License-Identifier: MIT
package types

import "encoding/json"

// AppendToInterface using the lazy approach.
//
// dst must be a pointer.
func AppendToInterface(src, dst interface{}) (err error) {
	var buffer []byte
	if buffer, err = json.Marshal(src); err != nil {
		return
	}
	err = json.Unmarshal(buffer, dst)

	return
}
