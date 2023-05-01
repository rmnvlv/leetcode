package Easy

import (
	"bytes"
	"strings"
)

// Runtime 1 ms Beats 66.67% Memory 1.9 MB Beats 39.92%
func defangIPaddr1(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

// Runtime 0 ms Beats 100% Memory 1.9 MB Beats 39.92%
func defangIPaddr2(address string) string {
	var bb bytes.Buffer
	for _, char := range address {
		if char == '.' {
			bb.WriteString("[.]")
		} else {
			bb.WriteRune(char)
		}
	}
	return bb.String()
}
