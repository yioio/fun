package funArray

//import (
//	"log"
//	)
/**
  *
  * array_push — 将一个或多个单元压入数组的末尾（入栈）
  * @params unit start_index  返回的数组的第一个索引值
  * @params uint num 插入元素的数量。 必须大于或等于0
  * @params interface{} value 用来填充的值
  *
  * @return
  */
func Array_push(arr *[]interface{}, args ...interface{}) {
	*arr = append(*arr, args...)
}
