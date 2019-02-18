package funArray

func Array_merge(arr1, arr2 []interface{}) []interface{} {

	firstArr := arr1
	secondArr := arr2
	length := len(arr2)

	for i:=0 ; i< length;i++ {
		firstArr = append(firstArr, secondArr[i])
	} 
	
	return firstArr
}

func Array_merge_string(arr1, arr2 []string) []string {

	firstArr := arr1
	secondArr := arr2
	length := len(arr2)

	for i:=0 ; i< length;i++ {
		firstArr = append(firstArr, secondArr[i])
	}
	
	return firstArr
}

//
// func Array_merge(arr1, arr2 []interface{}) []interface{} {

// 	firstArr := []interface{}{}
// 	secondArr := []interface{}{}
// 	length := 0

// 	length1 := len(arr1)
// 	length2 := len(arr2)

// 	if (length1 < length2) {
// 		firstArr = arr2
// 		secondArr = arr1
// 		length = length1
// 	} else {		
// 		firstArr = arr1
// 		secondArr = arr2
// 		length = length2
// 	}

// 	for i:=0 ; i< length;i++ {
// 		firstArr = append(firstArr, secondArr[i])
// 	}
	
// 	return firstArr
// }

// func Array_merge_string(arr1, arr2 []string) []string {

	
// 	firstArr := []string{}
// 	secondArr := []string{}
// 	length := 0

// 	length1 := len(arr1)
// 	length2 := len(arr2)

// 	if (length1 < length2) {
// 		firstArr = arr2
// 		secondArr = arr1
// 		length = length1
// 	} else {		
// 		firstArr = arr1
// 		secondArr = arr2
// 		length = length2
// 	}

// 	for i:=0 ; i< length;i++ {
// 		firstArr = append(firstArr, secondArr[i])
// 	}

// 	return firstArr
// }

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
