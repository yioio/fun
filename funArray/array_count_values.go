package funArray

import(
	// "reflect"
)

func Array_key_exists(needle string, haystack map[string]interface{}) bool {
	var isExist = false
	for key , _ := range haystack{
		if(needle == key) {
			isExist = true
		}
	}
	return isExist
}
