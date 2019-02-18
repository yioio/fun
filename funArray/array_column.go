package funArray

import(
	"fmt"
)

func Array_column(array []map[string]interface{}, column string) []interface{} {

	retArray := []interface{}{}
	
	for _, value := range array {
		retArray = append(retArray, value[column])
	}

	return retArray
}

func Array_column_retArrayString(array []map[string]interface{}, column string) []string {

	retArray := []string{}
	
	for _, value := range array {
		retArray = append(retArray, fmt.Sprintf("%v", value[column]))
	}

	return retArray
}

func Array_column_mapStringString_retArrayString(array []map[string]string, column string) []string {
	
	retArray := []string{}
	
	for _, value := range array {
		retArray = append(retArray, value[column])
	}

	return retArray
}
