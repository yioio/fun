package funString

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funString"
)

func Test_Chunk_split(t *testing.T) {

	Convey(" 将字符串分割成小块", t, func() {

		So(funString.Chunk_split("abcd", 1, ""), ShouldEqual, "a\r\nb\r\nc\r\nd\r\n")

	})
}