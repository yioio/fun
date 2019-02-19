package funDatatype 

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funFilesystem"
)

func Test_File_exits(t *testing.T) {
	
	Convey("将数组key转为大写：检查第一个元素的值，如果仍然是interface{}会报错", t, func() {
		ret1 := funFilesystem.File_exits("/go/src/tuan_access/tools/initRouters.go")
		ret2 := funFilesystem.File_exits("/go/src/tuan_access/tools/initRouters.goabc")
		
        So(ret1 == true, ShouldEqual, true)
        So(ret2 == false, ShouldEqual, true)
	})
}
