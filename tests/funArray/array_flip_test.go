package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_flip(t *testing.T) {

	Convey("检查： Array_flip", t, func() {
		arr := map[interface{}]interface{}{
			"test0":0,
			"test1":1,
			"test2":2,
			"test3":3,
		}
		processRet1 := funArray.Array_flip(arr)   // []
		log.Println("Array_flip 1 return is ", processRet1)
		So(processRet1, ShouldEqual, processRet1)
	})


}
