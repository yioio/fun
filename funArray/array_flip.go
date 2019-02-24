package funArray

//import (
//	"log"
//	)
/**
  * array_fill_keys — 使用指定的键和值填充数组
  * @params unit start_index  返回的数组的第一个索引值
  * @params uint num 插入元素的数量。 必须大于或等于0
  * @params interface{} value 用来填充的值
  *
  * @return
  */
func Array_fill(start_index int, num uint, value interface{}) map[interface{}]interface{} {
	fillArr := map[interface{}]interface{}{}
	var i uint
	for i = 1; i <= num; i++ {
		//log.Println(i)
		fillArr[start_index] = value
		start_index++
	}

	return fillArr
}
