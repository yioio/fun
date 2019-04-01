package funString

import "strings"

func Strpos(haystack, needle string, offset int) int {

	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)

	if pos == -1 {
		return -1
	}

	return pos + offset
}