package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	//"github.com/yioio/fun/funArray"
)


func Test_Array_chunk(t *testing.T) {

	Convey("check： Test_Array_chunk", t, func() {
		testArr := []interface{}{"test", "test2", "test3", 0,1,2,3,4,5,6,7}
		processRet1 := Array_chunk(testArr, 1)
		log.Println("Test_Array_chunk processRet1 return is ", processRet1)
		processRet2 := Array_chunk(testArr, 3)
		log.Println("Test_Array_chunk processRet2 return is ", processRet1)
		processRet3 := Array_chunk(testArr, 5)
		log.Println("Test_Array_chunk processRet3 return is ", processRet3)
		So(processRet1, ShouldEqual, true)
		So(processRet2, ShouldEqual, true)
		So(processRet3, ShouldEqual, false)
	})
}

