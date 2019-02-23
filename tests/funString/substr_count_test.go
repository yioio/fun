package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func SubstrCount_test(t *testing.T) {

	Convey("计算字串出现的次数", t, func() {

		So(funString.SubstrCount("abcabc", "bc"), ShouldEqual, 2)

	})
}