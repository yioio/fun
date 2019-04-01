package funTime

import (
	"time"
	"fmt"
)

func init() {
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
}

//例如：2019-01-31 00:00:00
func Strtotime(strtime interface{}) int64 {

	return Strtotime_DateTime(strtime)
}

//例如：2019-01-31 00:00:00
func Strtotime_DateTime(strtime interface{}) int64 {
	
	return strToTime_DateTime(strtime)
}

//例如：2019-01-31
func Strtotime_Date(strtime interface{}) int64 {

	return strToTime_DateTime(fmt.Sprintf("%v 00:00:00", strtime))
}

//外部不可以直接使用， 例如：2019-01-31 00:00:00 
func strToTime_DateTime(strtimeI interface{}) int64 {

	strtime := fmt.Sprintf("%v", strtimeI)

	time1, err := time.ParseInLocation(FORMAT_DATETIME, strtime, cstSh)
	
	if (err != nil) {

		return 0
	}

	return time1.Unix()
}
