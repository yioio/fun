package funArray

//import (
//	"log"
//	)
/**
  * array_reverse — 返回单元顺序相反的数组
  * @params []interface{} reverseArr 要循环的数组
  *
  * @return  reverseArr
  */
func Array_reverse( reverseArr []interface{}) []interface{}{
	for i, j := 0, len(reverseArr)-1; i < j; i, j = i+1, j-1 {
		reverseArr[i], reverseArr[j] = reverseArr[j], reverseArr[i]
	}
	return reverseArr
}

