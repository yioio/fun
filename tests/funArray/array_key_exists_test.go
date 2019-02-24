package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	//"github.com/yioio/fun/funArray"
)


//func Test_Array_key_exists(t *testing.T) {
//
//	Convey("检查： Array_column", t, func() {
//		toMaps := []interface{}{"123", "234"}
//		processRet1 := funArray.In_array("123", toMaps)
//		processRet2 := funArray.In_array("234", toMaps)
//		processRet3 := funArray.In_array("", toMaps)
//		processRet4 := funArray.In_array("890", toMaps)
//		So(processRet1, ShouldEqual, true)
//		So(processRet2, ShouldEqual, true)
//		So(processRet3, ShouldEqual, false)
//		So(processRet4, ShouldEqual, false)
//	})
//}


func Test_Array_key_exists(t *testing.T) {

	Convey("检查： Array_key_exists", t, func() {
		toMaps := map[string]interface{}{"123":"123", "234":"234"}
		//toMaps := []string{"123", "234"}
		processRet1 := Array_key_exists("123", toMaps)
		log.Println("processRet1返回结果是", processRet1)
		processRet2 := Array_key_exists("234", toMaps)
		processRet3 := Array_key_exists("", toMaps)
		log.Println("processRet3 返回结果是", processRet3)
		So(processRet1, ShouldEqual, true)
		So(processRet2, ShouldEqual, true)
		So(processRet3, ShouldEqual, false)
	})
}

