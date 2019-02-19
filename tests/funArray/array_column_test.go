package funArray

import(
	"reflect"
)

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

func Test_Array_column(t *testing.T) {
	
	Convey("检查： Array_column", t, func() {
		processRet := funArray.Array_column([]map[string]interface{}{{"apple":"123","b":"123"}, {"apple":"124","b":"123"},{"apple":"125","b":"123"}}, string("apple"))
		ret := []interface{}{"123",  "124", "125"}
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(ret))
        So(processRet[0], ShouldEqual, ret[0])
        So(processRet[1], ShouldEqual, ret[1])
        So(processRet[2], ShouldEqual, ret[2])
	})
	Convey("检查： Array_column_retArrayString", t, func() {
		processRet := funArray.Array_column_retArrayString([]map[string]interface{}{{"apple":"123","b":"123"}, {"apple":"124","b":"123"},{"apple":"125","b":"123"}}, string("apple"))
		ret := []string{"123",  "124", "125"}
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(ret))
        So(processRet[0], ShouldEqual, ret[0])
        So(processRet[1], ShouldEqual, ret[1])
        So(processRet[2], ShouldEqual, ret[2])
	})
	Convey("检查： Array_column_mapStringString_retArrayString", t, func() {
		processRet := funArray.Array_column_mapStringString_retArrayString([]map[string]string{{"apple":"123","b":"123"}, {"apple":"124","b":"123"},{"apple":"125","b":"123"}}, string("apple"))
		ret := []string{"123",  "124", "125"}
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(ret))
        So(processRet[0], ShouldEqual, ret[0])
        So(processRet[1], ShouldEqual, ret[1])
        So(processRet[2], ShouldEqual, ret[2])
	}) 
}
