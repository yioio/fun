package funArray

import(
	// "reflect"
)

func Array_search_string(needle string, haystack []string) int {
	length := len(haystack)
	 
	for i:= 0;i< length;i++ {
		if (haystack[i] == needle) {
			return i
		}
	}

	return -1
}

func Array_search(needle interface{}, haystack []interface{}) int {

	length := len(haystack)

	for i:= 0; i<length; i++ {
		if (haystack[i] == needle) {
			return i
		}
	}

	return -1
}
