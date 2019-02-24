package funFilesystem

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funFilesystem"
)

func Test_Basename(t *testing.T) {

	Convey("给出一个包含有指向一个文件的全路径的字符串，本函数返回基本的文件名", t, func() {
		So(funFilesystem.Basename("/etc/sudoers.d", ".d"), ShouldEqual, "sudoers")
		So(funFilesystem.Basename("/etc/passwd", ""), ShouldEqual, "passwd")
		So(funFilesystem.Basename("/etc/", ""), ShouldEqual, "etc")
	})
}