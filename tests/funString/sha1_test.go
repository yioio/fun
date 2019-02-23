package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Sha1_test(t *testing.T) {

	Convey("计算字符串的 sha1 散列值", t, func() {

		So(funString.Sha1("apple"), ShouldEqual, "d0be2dc421be4fcd0172e5afceea3970e2f3d940")

	})
}