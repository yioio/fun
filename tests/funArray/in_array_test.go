package funArray

import(
	// "reflect"
	// "sort"
)

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

func Test_In_array_string(t *testing.T) {
	
	Convey("检查： Array_column", t, func() {
		toMaps := []string{"123", "234"}
		processRet1 := funArray.In_array_string("123", toMaps)
		processRet2 := funArray.In_array_string("234", toMaps)
		processRet3 := funArray.In_array_string("", toMaps)
		processRet4 := funArray.In_array_string("890", toMaps)
        So(processRet1, ShouldEqual, true)
        So(processRet2, ShouldEqual, true)
        So(processRet3, ShouldEqual, false)
        So(processRet4, ShouldEqual, false)
	})	
}

func Test_In_array(t *testing.T) {
	
	Convey("检查： Array_column", t, func() {
		toMaps := []interface{}{"123", "234"}
		processRet1 := funArray.In_array("123", toMaps)
		processRet2 := funArray.In_array("234", toMaps)
		processRet3 := funArray.In_array("", toMaps)
		processRet4 := funArray.In_array("890", toMaps)
        So(processRet1, ShouldEqual, true)
        So(processRet2, ShouldEqual, true)
        So(processRet3, ShouldEqual, false)
        So(processRet4, ShouldEqual, false)
	})	
}
