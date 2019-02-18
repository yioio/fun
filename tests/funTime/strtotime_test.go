package funTime

import (
    "testing" 
	. "github.com/smartystreets/goconvey/convey" 
	"go.increase/fun/funTime"
	// "fmt"
	"time"
)

func Test_Strtotime(t *testing.T) {

	Convey("Test Strtotime", t, func() {

		//时区正确下的比较
		cstSh, _ := time.LoadLocation("Asia/Shanghai") //上海 
		date, _ := time.ParseInLocation(funTime.FORMAT_DATETIME, "2017-05-11 00:00:00", cstSh)
		time1 := funTime.Strtotime("2017-05-11 00:00:00")
		time2 := funTime.Strtotime_DateTime("2017-05-11 00:00:00")
		So(time1 > 0, ShouldBeTrue)
		So(time1, ShouldEqual, date.Unix())
		So(time1, ShouldEqual, time2)

		//晚8个时区
		date2, _ := time.Parse(funTime.FORMAT_DATETIME, "2017-05-11 00:00:00")
		So(date2.Unix() - time1, ShouldEqual, 8*3600)
		
		//测试日期
		time3 := funTime.Strtotime_Date("2017-05-11")

		date3, _ := time.ParseInLocation(funTime.FORMAT_DATETIME, "2017-05-11 00:00:00", cstSh)
		So(date3.Unix(), ShouldEqual, time3)
	})
}
