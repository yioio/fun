package funTime

import (
	"time"
	"log"
)

func init() {
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
}

//例如：2019-01-31 00:00:00
func Strtotime(strtime string) int64 {

	return Strtotime_DateTime(strtime)
}

//例如：2019-01-31 00:00:00
func Strtotime_DateTime(strtime string) int64 {
	
	return strToTime_DateTime(strtime)
}

//例如：2019-01-31
func Strtotime_Date(strtime string) int64 {

	return strToTime_DateTime(strtime + " 00:00:00") 
}

//外部不可以直接使用， 例如：2019-01-31 00:00:00 
func strToTime_DateTime(strtime string) int64 {

	time1, err := time.ParseInLocation(FORMAT_DATETIME, strtime, cstSh)
	
	if (err != nil) {

		log.Println(err)

		return 0
	}

	return time1.Unix()
}
