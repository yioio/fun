package funArray

import(
	"reflect"
)

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

func Test_Array_Merge(t *testing.T) {
	
	Convey("检查： Array_merge", t, func() {

		arr1 := []interface{}{"a","3","c"}
		arr2 := []interface{}{"2","b"}

		processRet := funArray.Array_merge(arr1, arr2)
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(arr1))
        So(funArray.Array_search("a", processRet), ShouldEqual, 0)
        So(funArray.Array_search("2", processRet), ShouldEqual, 3)
        So(funArray.Array_search("abc", processRet), ShouldEqual, -1)
	})

	Convey("检查： Array_merge_string", t, func() {

		arr1 := []string{"1","a","3","c"}
		arr2 := []string{"2","b"}

		processRet := funArray.Array_merge_string(arr1, arr2)
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(arr1))
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(arr1))
        So(funArray.Array_search_string("a", processRet), ShouldEqual, 1)
        So(funArray.Array_search_string("2", processRet), ShouldEqual, 4)
        So(funArray.Array_search_string("abc", processRet), ShouldEqual, -1)
	})
	
	Convey("检查： Map_merge", t, func() {

		arr1 := map[string]interface{}{"1":"a","3":"c"}
		arr2 := map[string]interface{}{"2":"b"}

		processRet := funArray.Map_merge(arr1, arr2)
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(arr1))
        So(processRet["1"], ShouldEqual, interface{}("a"))	
        So(processRet["2"], ShouldEqual, interface{}("b"))	
        So(processRet["3"], ShouldEqual, interface{}("c"))	
	}) 

	Convey("检查： Map_merge_string", t, func() {

		arr1 := map[string]string{"1":"a","3":"c"}
		arr2 := map[string]string{"2":"b"}

		processRet := funArray.Map_merge_string(arr1, arr2)
        So(reflect.TypeOf(processRet), ShouldEqual, reflect.TypeOf(arr1))
        So(processRet["1"], ShouldEqual, string("a"))	
        So(processRet["2"], ShouldEqual, string("b"))	
        So(processRet["3"], ShouldEqual, string("c"))	
	})
}
