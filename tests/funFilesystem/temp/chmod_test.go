package temp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/yioio/fun/funFilesystem"
)

func Test_chmod(t *testing.T)  {
	Convey("修改文件权限模式", t, func() {
		res := funFilesystem.Chmod("C:\\Users\\chenhongsheng\\Desktop\\fun\\doc\\zh\\test.md", 0600)

		So(res, ShouldEqual, true)
	})
}
