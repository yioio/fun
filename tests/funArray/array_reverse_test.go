package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_reverse(t *testing.T) {

	Convey("检查： Array_reverse", t, func() {
		Arrayreverse := []interface{}{"test", 123,"123468", []interface{}{"sfasfdsa", 123}}
		processRet1 := funArray.Array_reverse(Arrayreverse)   // []
		log.Println("Array_reverse 1 return is ", processRet1)
		//So(processRet1, ShouldEqual, processRet1)
	})


}
