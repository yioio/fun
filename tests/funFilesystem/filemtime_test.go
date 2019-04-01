package funFilesystem

import (
    "testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yioio/fun/funTime"
	"github.com/yioio/fun/funFilesystem"
)

func Test_Filemtime(t *testing.T) {
	
	Convey("将数组key转为大写：检查第一个元素的值，如果仍然是interface{}会报错", t, func() {
		processRet := funFilesystem.Filemtime("/go/src/tuan_access/tools/initRouters.go")
		earlyTime := funTime.Strtotime_Date("2019-02-01")
        So(processRet > 0, ShouldEqual, true)
        So(processRet < funTime.Time(), ShouldEqual, true)
        So(processRet > earlyTime, ShouldEqual, true)
	})
}
