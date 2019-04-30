package funString

import (
	"github.com/yioio/fun/funString"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Join(t *testing.T) {

	Convey("将一个一维数组的值转化为字符串", t, func() {

		So(funString.Join("|", []string{"abc","aaa"}), ShouldEqual, "abc|aaa")

	})
}
