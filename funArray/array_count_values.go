package funArray

import(
	// "reflect"
)

func Array_count_values(arr []interface{}) map[interface{}]uint {
	countValues := make(map[interface{}]uint)
	for _, v := range arr {
		if c, ok := countValues[v]; ok {
			countValues[v] = c + 1
		} else {
			countValues[v] = 1
		}
	}

	return countValues
}
