package funDatatype 

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funDatatype"
	"reflect"
	"log"
)

func Test_Array_interface_to_array_string(t *testing.T) {
	
	Convey("将数组key转为大写：检查第一个元素的值，如果仍然是interface{}会报错", t, func() {
		processRet := funDatatype.ArrayInterface_to_arrayString([]interface{}{"123"})
        So(processRet[0], ShouldEqual, []string{"123"}[0])
	})

	Convey("将数组key转为大写:检查返回类型", t, func() {
		processRet := funDatatype.ArrayInterface_to_arrayString([]interface{}{"123"})
        So(reflect.TypeOf(processRet[0]), ShouldEqual, reflect.TypeOf(string("123")))
	})	
}


/**
// 将  []string 转成 []interface{}
测试用例
 */
func Test_StringArr_to_interfaceArr(t *testing.T) {

	Convey("检查： Test_StringArr_to_interfaceArr", t, func() {
		orgArr := []string{"我是测试a", "我是测试b"}
		log.Println("原来的数据类型是", reflect.TypeOf(orgArr))
		log.Println("Test_StringArr_to_interfaceArr last", orgArr)

		res := funDatatype.StringArr_to_interfaceArr(orgArr)   // []
		log.Println("转换之后的的数据类型是", reflect.TypeOf(res))
		log.Println("Test_StringArr_to_interfaceArr 1 return is ", res)
	})



}
