package funString 

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Md5(t *testing.T) {
	
	Convey("默认 md5", t, func() {
		
        So(funString.Md5(""), ShouldEqual, "d41d8cd98f00b204e9800998ecf8427e")
        So(funString.Md5("abc123"), ShouldEqual, "e99a18c428cb38d5f260853678922e03")
        So(funString.Md5("1234567890"), ShouldEqual, "e807f1fcf82d132f9bb018ca6738a19f")
        So(funString.Md5("ka-)((*&^"), ShouldEqual, "be4b2a8cb624dcf8540ad196707443b2")
	})
}
