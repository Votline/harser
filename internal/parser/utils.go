// Package parser utils.go contains help-funcs for parser package
package parser

import (
	"bytes"
)

// rangeByByte call yield for each matching row
func rangeByByte(src []byte, sep byte, yield func(start, end int)) {
	start := 0
	for start < len(src) {
		end := bytes.IndexByte(src[start:], sep)
		if end == -1 {
			end = len(src)
		} else {
			end += start
		}
		yield(start, end)
		start = end + 1
	}
}

// trimSpaceBytes trim spaces from slice byte
// modifies slice, no allocations
func trimSpaceBytes(buf *[]byte) {
	src := *buf
	start := 0
	for start < len(src) && isSpaceByte(src[start]) {
		start++
	}
	end := len(src)
	for end > start && isSpaceByte(src[end-1]) {
		end--
	}
	*buf = src[start:end]
}

// isSpaceByte returns true if byte is space
func isSpaceByte(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}
