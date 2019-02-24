package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_fill_keys(t *testing.T) {

	Convey("检查： Array_fill_keys", t, func() {
		t1_kyes := []interface{}{"test",123,"0123"}
		t1_values := []interface{}{"test2"}
		processRet1 := funArray.Array_fill_keys(t1_kyes, t1_values)   // []
		log.Println("Test_Array_fill_keys 1 return is ", processRet1)
		So(processRet1, ShouldEqual, processRet1)
	})
}
