package funDatatype

import(
	"fmt"
)

//将 map[string]interface{} 转为 map[string]string
func Map_string_interface_to_string_string(maps map[string]interface{}) map[string]string {

	mapsRet := map[string]string{}
	for key, value := range maps {
		mapsRet[key] = fmt.Sprintf("%v", value)
	}

	return mapsRet
}


//将 map[string]int 转为 map[string]string
func Map_string_int_to_string_string(maps map[string]int) map[string]string {

	mapsRet := map[string]string{}
	for key, value := range maps {
		mapsRet[key] = fmt.Sprintf("%v", value)
	}

	return mapsRet
}

//将 map[string]float64 转为 map[string]string
func Map_string_float64_to_string_string(maps map[string]float64) map[string]string {

	mapsRet := map[string]string{}
	for key, value := range maps {
		mapsRet[key] = fmt.Sprintf("%v", value)
	}

	return mapsRet
}

//将 map[string]string  转为 map[string]interface{}
func Map_string_string_to_interface_string(maps map[string]string ) map[string]interface{} {

	mapsRet := make(map[string]interface{})
	for key, value := range maps {
		mapsRet[key] = fmt.Sprintf("%v", value)
	}

	return mapsRet
}
