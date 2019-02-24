package funArray

import(
	// "reflect"
)

func Array_diff(array1 , array2  []interface{})  []interface{} {
	diffArr :=  []interface{}{}
	for _, val := range array1 {
		if In_array(val, array2) == false {
			diffArr = append(diffArr, val)
		}
	}
		return diffArr
}
