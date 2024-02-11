// Copyright 2023 Jens Weißkopf. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
// Forked, changed and extend from https://github.com/JacThomp/docstract
// Special thx to JacThomp

package docstract

// StripSeperators removes all the random 0 bytes
func StripSeperators(s string) string {
	iBytes := []byte(s)
	oBytes := []byte{}

	if len(iBytes) >= 3 {
		offset := 0
		if iBytes[0] == iBytes[2] && iBytes[0] == byte(0) {
			offset = 1
		}

		for i := offset; i < len(iBytes); i += 2 {
			oBytes = append(oBytes, iBytes[i])
		}
	}

	return string(oBytes)
}
