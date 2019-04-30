package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Html_entity_decode(t *testing.T) {

	Convey("使用反斜线引用字符串", t, func() {

		So(funString.Html_entity_decode("I'll &quot;walk&quot; the &lt;b&gt;dog&lt;/b&gt; now"), ShouldEqual, "I'll \"walk\" the <b>dog</b> now")

	})
}