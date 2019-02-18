package funArray

import(
	// "reflect"
	// "sort"
)

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"go.increase/fun/funArray"
)

func Test_Array_search_string(t *testing.T) {
	
	Convey("检查： Array_column", t, func() {

		toMaps := []string{"123", "234"}

		processRet1 := funArray.Array_search_string("123", toMaps)
		processRet2 := funArray.Array_search_string("234", toMaps)
		processRet3 := funArray.Array_search_string("", toMaps)
		processRet4 := funArray.Array_search_string("890", toMaps)

        So(processRet1, ShouldEqual, 0)
        So(processRet2, ShouldEqual, 1)
        So(processRet3, ShouldEqual, -1)
        So(processRet4, ShouldEqual, -1)
	})	
}

func Test_Array_search(t *testing.T) {
	
	Convey("检查： Array_column", t, func() {
		toMaps := []interface{}{"123", "234"}
		processRet1 := funArray.Array_search("123", toMaps)
		processRet2 := funArray.Array_search("234", toMaps)
		processRet3 := funArray.Array_search("", toMaps)
		processRet4 := funArray.Array_search("890", toMaps)
        So(processRet1, ShouldEqual, 0)
        So(processRet2, ShouldEqual, 1)
        So(processRet3, ShouldEqual, -1)
        So(processRet4, ShouldEqual, -1)
	})	
}
