package funArray

import (
	"testing"
	"log"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

//
func Test_Array_pop(t *testing.T) {

	Convey("检查： Array_pop", t, func() {
		orgArr := []interface{}{"test", 000, "ok"}
		popElem := funArray.Array_pop(&orgArr)   // []
		log.Println("Array_pop 1 return is ", orgArr, popElem)
		//So(orgArr, ShouldEqual, orgArr)
	})


}
