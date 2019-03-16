package funDatatype 

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funDatatype"
	"reflect"
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
