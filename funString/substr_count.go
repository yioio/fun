package funString

import "strings"

func SubstrCount(haystack string ,  needle string ) int {

	return strings.Count(haystack, needle)
}