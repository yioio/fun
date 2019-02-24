package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_push(t *testing.T) {

	Convey("检查： Array_push", t, func() {
		orgArr := []interface{}{"test", 000, "ok"}
		funArray.Array_push(&orgArr, "test0")   // []
		log.Println("Test_Array_push 1 return is ", orgArr)
		//So(orgArr, ShouldEqual, orgArr)
	})


}
