package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Htmlentities(t *testing.T) {

	Convey("使用反斜线引用字符串", t, func() {

		So(funString.Htmlentities("A 'quote' is <b>bold</b>"), ShouldEqual, "A &#39;quote&#39; is &lt;b&gt;bold&lt;/b&gt;")

	})
}