package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_key_first(t *testing.T) {

	Convey("检查： Array_key_first", t, func() {
		orgArr := map[interface{}]interface{}{"test":"00", 0:000, 1:"ok"}
		res := funArray.Array_key_first(orgArr)   // []
		log.Println("Array_key_first 1 return is ", res)
		So(res, ShouldEqual, res)
	})


}
