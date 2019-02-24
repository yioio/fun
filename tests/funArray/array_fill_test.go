package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_fill_keys(t *testing.T) {

	Convey("检查： Array_fill", t, func() {
		processRet1 := funArray.Array_fill(5, 8, "123456")   // []
		log.Println("Test_Array_fill 1 return is ", processRet1)
		So(processRet1, ShouldEqual, processRet1)
	})


}
