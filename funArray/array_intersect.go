package funArray

import(
	// "reflect"
)


func Array_intersect(array1, array2 []interface{}) []interface{} {
	diffArr :=  []interface{}{}
	for _, val := range array1 {
		if In_array(val, array2) == true {
			diffArr = append(diffArr, val)
		}
	}
	return diffArr
}
