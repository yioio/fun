package funDatatype

import(
	// "reflect"
	"fmt"
)

// []interface{} to []string
func ArrayInterface_to_arrayString(array []interface{}) []string {

	arrayRet := []string{}
	
	for _, value := range array {
		arrayRet = append(arrayRet, fmt.Sprintf("%v", value))
	}
	
	return arrayRet
}

// []int to []string
func ArrayInt_to_arrayString(array []int) []string {

	arrayRet := []string{}
	
	for _, value := range array {
		arrayRet = append(arrayRet, fmt.Sprintf("%v", value))
	}
	
	return arrayRet
}

// []float64 to []string
func ArrayFloat64_to_arrayString(array []int) []string {
	
	arrayRet := []string{}
	
	for _, value := range array {
		arrayRet = append(arrayRet, fmt.Sprintf("%v", value))
	}
	
	return arrayRet
}

// 将  []string 转成 []interface{}
func StringArr_to_interfaceArr(interfaceArr []string) ([]interface{}) {
	var strArr []interface{}
	for _, v := range interfaceArr {
		strArr = append(strArr, v)
	}
	return strArr
}
