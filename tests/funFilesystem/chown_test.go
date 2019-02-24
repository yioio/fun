package funFilesystem

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
	"github.com/yioio/fun/funFilesystem"
)

func Test_Chown(t *testing.T)  {
	Convey("修改文件所有者", t, func() {
		guid := os.Getuid()
		uid := os.Getuid()
		res := funFilesystem.Chown("C:\\Users\\chenhongsheng\\Desktop\\fun\\doc\\zh\\test.md", guid, uid)

		So(res, ShouldEqual, true)
	})
}
