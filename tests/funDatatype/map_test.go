package funDatatype 

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funDatatype"
	"reflect"
	"log"
)



/**
将 map[string]string  转为 map[string]interface{}
测试用例
 */
func Test_Map_string_string_to_interface_string(t *testing.T) {

	Convey("检查： Test_Map_string_string_to_interface_string", t, func() {
		orgArr := map[string]string{"a":"我是测试a", "b":"我是测试b"}
		log.Println("原来的数据类型是", reflect.TypeOf(orgArr))
		log.Println("Test_Map_string_string_to_interface_string last", orgArr)

		res := funDatatype.Map_string_string_to_interface_string(orgArr)   // []
		log.Println("转换之后的的数据类型是", reflect.TypeOf(res))
		log.Println("Test_Map_string_string_to_interface_string 1 return is ", res, orgArr)
	})


}
