package funArray

//import(
//	"github.com/yioio/fun/funVar"
//	"math"
//)

/**
 * array_pop — 弹出数组最后一个单元（出栈）
 * @param *[]interface{} s 数组的地址
 *
 * @return interface{} 被弹出的数据
 */
func Array_pop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

