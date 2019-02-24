package funArray

//import (
//	"log"
//	)
/**
  * array_flip — 交换数组中的键和值
  * @params map[interface{}]interface{} arrayMap  原数据
  * @return
  */
func Array_flip(arrayMap map[interface{}]interface{}) map[interface{}]interface{} {
	fillArr := map[interface{}]interface{}{}
	for i, v := range arrayMap {
		fillArr[v] = i
	}
	return fillArr
}
