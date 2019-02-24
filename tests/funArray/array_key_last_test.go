package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_key_last(t *testing.T) {

	Convey("检查： Array_key_last", t, func() {
		orgArr := map[interface{}]interface{}{"test":"00", 0:000, 1:"ok"}
		log.Println("orgArrorgArrorgArr last", orgArr)
		res := funArray.Array_key_last(orgArr)   // []
		log.Println("Array_key_last 1 return is ", res, orgArr)
		So(res, ShouldEqual, res)
	})


}
