package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Explode(t *testing.T) {

	Convey("使用一个字符串分割另一个字符串", t, func() {
		ret := funString.Explode("|", "abc|bcd|aaa")
		So(ret[0], ShouldEqual, []string{"abc","bcd","aaa"}[0])
		So(ret[1], ShouldEqual, []string{"abc","bcd","aaa"}[1])
		So(ret[2], ShouldEqual, []string{"abc","bcd","aaa"}[2])
	})
}