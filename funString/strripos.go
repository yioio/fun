package funString

import "strings"

func Strripos(haystack, needle string, offset int) int {
	pos, length := 0, len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		haystack = haystack[:offset+length+1]
	} else {
		haystack = haystack[offset:]
	}

	pos = strings.LastIndex(strings.ToLower(haystack), strings.ToLower(needle))

	if offset > 0 && pos != -1 {
		pos += offset
	}
	
	return pos
}