package funArray

import(
	// "reflect"
)

func Array_key_exists(needle interface{}, haystack map[interface{}]interface{}) bool {
	var isExist = false
	_, isExist = haystack[needle]
	return isExist
}
