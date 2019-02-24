package funArray

import(
	// "fmt"
	// "reflect"
	// "strconv"
	"sort"
)

func Array_keys(array map[string]interface{}) []string {

	retArray := []string{}
	for key,_ := range array {
		retArray = append(retArray, key)
	}
	sort.Strings(retArray)

	return retArray
}

//key的类型为int
func Array_keys_int(array map[int]interface{}) []int {

	retArray := []int{}
	for key,_ := range array {
		retArray = append(retArray, key)
	}
	sort.Ints(retArray)

	return retArray
}

//key的类型为interface{}
func Array_keys_interface(array map[interface{}]interface{}) []interface{} {

	retArray := []interface{}{}
	for key,_ := range array {
		retArray = append(retArray, key)
	}
 
	// sort.Sort(retArray)

	return retArray
}
