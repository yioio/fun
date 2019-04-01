package funArray

import(
	// "reflect"
)

/*
	不建议在实际程序中使用本函数，建议使用：
	if _, exist := haystack[needle]; exist {

	}
	if _, exist := haystack[needle]; !exist {

	}
*/

func Array_key_exists(needle interface{}, haystack map[interface{}]interface{}) bool {
	var isExist = false
	_, isExist = haystack[needle]
	return isExist
}

func Array_key_exists_string(needle string, haystack map[string]interface{}) bool {
	var isExist = false
	_, isExist = haystack[needle]
	return isExist
}

