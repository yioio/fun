package funArray

import(
	"reflect"
	"sort"
)

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"go.increase/fun/funArray"
)

func Test_Array_keys(t *testing.T) {
	
	Convey("检查： Array_column", t, func() {
		processRet := funArray.Array_keys(map[string]interface{}{"apple":"123","b":"123"})
		ret := []string{"apple", "b"}
		sort.Strings(ret)
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(ret))
        So(processRet[0], ShouldEqual, ret[0])
        So(processRet[1], ShouldEqual, ret[1])
	})
	
	Convey("检查： Array_column_int", t, func() {
		processRet := funArray.Array_keys_int(map[int]interface{}{1:"123",2:"123"})
		ret := []int{int(1), int(2)}
		sort.Ints(ret)
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(ret))
        So(processRet[0], ShouldEqual, ret[0])
        So(processRet[1], ShouldEqual, ret[1]) 
	})
	
	Convey("检查： Array_column_interface", t, func() { 
		processRet := funArray.Array_keys_interface(map[interface{}]interface{}{1:"123","apple":"123"})
		ret := []interface{}{"apple", 1}
		// sort.Sort(ret)
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(ret))
        So(funArray.In_array(1, processRet), ShouldEqual, funArray.In_array(1, ret))
        So(funArray.In_array("apple", processRet), ShouldEqual, funArray.In_array("apple", ret))
	})
}
