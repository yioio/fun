package funArray

//对于切片不建议使用本方法，而要直接使用 append
func Array_merge(arr1, arr2 []interface{}) []interface{} {

	return append(arr1, arr2...)
}

//对于切片不建议使用本方法，而要直接使用 append
func Array_merge_string(arr1, arr2 []string) []string {
	return append(arr1, arr2...)
}

func Map_merge(arr1, arr2 map[string]interface{}) map[string]interface{} {

	newArr := arr1

	for key, value := range arr2 {
		newArr[key] = value
	}

	return newArr
}

func Map_merge_string(arr1, arr2 map[string]string) map[string]string {
	
	newArr := arr1

	for key, value := range arr2 {
		newArr[key] = value
	}

	return newArr
}

func Map_merge_keyInterface(arr1, arr2 map[interface{}]interface{}) map[interface{}]interface{} {

	newArr := arr1

	for key, value := range arr2 {
		newArr[key] = value
	}

	return newArr
}
