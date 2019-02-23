package funString

import "strings"

func Stristr(haystack string, needle string, beforeNeedle bool) string {

	if needle == "" {
		return ""
	}

	idx := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))


	if idx == -1 {
		return ""
	}

	if beforeNeedle {

		return haystack[:idx]

	} else {

		return haystack[idx:]
	}

}