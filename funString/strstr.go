package funString

import "strings"



func Strstr(haystack string, needle string, beforeNeedle bool) string {

	if needle == "" {
		return ""
	}

	idx := strings.Index(haystack, needle)


	if idx == -1 {
		return ""
	}

	if beforeNeedle {

		return haystack[:idx]

	} else {

		return haystack[idx:]
	}

}