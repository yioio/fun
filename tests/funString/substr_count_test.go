package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_SubstrCount(t *testing.T) {

	Convey("计算字串出现的次数", t, func() {

		So(funString.SubstrCount("abcabc", "bc"), ShouldEqual, 2)

	})
}