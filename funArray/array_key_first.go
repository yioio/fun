package funArray

import(
	// "reflect"
)
/**
 * array_key_first — Gets the first key of an array
 * @params map[interface{}]interface{} array An array.
 *
 * @return interface{}
 */
func Array_key_first(array map[interface{}]interface{}) interface{} {
	var k interface{}
	for key, _ := range array {
		k = key
		break;
	}
	return k
}
