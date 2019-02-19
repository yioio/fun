package funArray

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funArray"
)

func Test_Array_change_key_case(t *testing.T) {
	
	Convey("将数组key转为大写", t, func() {
		processRet := funArray.Array_change_key_case(map[string]interface{}{"teST":"123"}, funArray.CASE_UPPER)
        So(processRet["TEST"].(string), ShouldEqual, "123")
	})
	
	Convey("将数组key转为小写", t, func() {
		processRet := funArray.Array_change_key_case(map[string]interface{}{"teST":"123"}, funArray.CASE_LOWER)
        So(processRet["test"].(string), ShouldEqual, "123")
	})

}
