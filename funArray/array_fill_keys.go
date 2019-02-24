package funArray

/**
  * array_fill_keys — 使用指定的键和值填充数组
  * @params []interface{} arrayKeys  使用该数组的值作为键
  * @params []interface{} arrayValues  填充使用的值 这里是一个数组，方便支持多个值
  */
func Array_fill_keys(arrayKeys [] interface{}, arrayValues []interface{}) map[interface{}]interface{} {
	fillArr := map[interface{}]interface{}{}
	for _, v := range arrayKeys {
		for _, val := range arrayValues {
			fillArr[v] = val
		}
	}
	return fillArr
}

