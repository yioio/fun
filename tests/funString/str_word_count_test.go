package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Str_word_count(t *testing.T) {

	Convey("返回字符串中单词的使用情况", t, func() {

		So(funString.Str_word_count("Hello friend you're looking good today"), ShouldEqual, []string{"Hello",
		"friend","you're","looking","good","today"})

	})
}