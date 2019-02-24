package funTime

import (
    "testing" 
	. "github.com/smartystreets/goconvey/convey" 
	"github.com/yioio/fun/funTime"
	"github.com/yioio/fun/funVar"
	"strings"
	"os/exec"
	// "fmt"
	"time"
)

func Test_Datetime(t *testing.T) {

	Convey("Test Date", t, func() {

		cmd := exec.Command("date", "+%s")
		out, _ := cmd.Output()
		newStr:= strings.Trim(string(out), "\n")
		shelltime := funVar.Int64val(newStr)

		//时区正确下的比较
		timestamp := funTime.Strtotime_Date("2017-05-11")
		date := funTime.Datetime(timestamp)
		So(date, ShouldEqual, "2017-05-11 00:00:00")
		
		timestamp = time.Now().Unix()
		date = funTime.Datetime(timestamp) 
		today := funTime.Datetime_today() 
		tomorrow := funTime.Datetime_tomorrow() 
		So(date, ShouldEqual, funTime.Datetime(shelltime))
		So(funTime.Date_now() , ShouldEqual, funTime.Date_today()) 

		So(funTime.Strtotime_DateTime(today), ShouldEqual, timestamp)
		So(shelltime, ShouldEqual, timestamp)
		So(funTime.Strtotime_DateTime(tomorrow) - 86400, ShouldEqual, shelltime)
		So(tomorrow, ShouldEqual, funTime.Datetime_daysafter(1, timestamp))
		So(today, ShouldEqual, funTime.Datetime_daysafter(0, timestamp))
		So(today, ShouldEqual, funTime.Datetime_daysbefore(1, timestamp + 86400))
	})
}
