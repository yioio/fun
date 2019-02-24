package funArray

import(
	// "reflect"
)
//type compArrayType = []interface{}

func Array_diff(array1  []interface{}, array2  []interface{})  []interface{} {
	diffArr :=  []interface{}{}
	for _, val := range array1 {
		if In_array(val, array2) == false {
			diffArr = append(diffArr, val)
		}
	}
		return diffArr
}
