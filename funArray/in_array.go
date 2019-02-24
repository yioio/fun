package funArray

import(
	// "reflect"
)

func In_array_string(needle string, haystack []string) bool {
	length := len(haystack)
	for i:= 0;i< length;i++ {
		if (haystack[i] == needle) {
			return true
		}
	}
	
	return false
}

func In_array(needle interface{}, haystack []interface{}) bool {

	length := len(haystack)
	for i:= 0;i< length;i++ {
		if (haystack[i] == needle) {
			return true
		}
	}

	return false
}
