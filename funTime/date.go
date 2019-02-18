package funTime

import (
	"time"
	// "log"
)

func init() {
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
}

//dateTimeFormat 格式如："2006-01-02 15:04:05" 
func Date(timestamp int64) string {
	return time.Unix(timestamp, 0).In(cstSh).Format(FORMAT_DATE)
}

//当前日期
func Date_now() string {
	return Date_today()
}

//当前日期
func Date_today() string {
	return time.Now().In(cstSh).Format(FORMAT_DATE)
}

//明天日期
func Date_tomorrow() string {
	return time.Unix(Time() + 86400, 0).In(cstSh).Format(FORMAT_DATE)
}

//多少天后
func Date_daysafter(days int64, timestamp int64) string {
	return Date_after(86400 * days, timestamp)
}

//多少天前
func Date_daysbefore(days int64, timestamp int64) string {
	return Date_before(86400 * days, timestamp)
}

//多少小时后
func Date_hoursafter(hours int64, timestamp int64) string {
	return Date_after(3600 * hours, timestamp)
}

//多少小时前
func Date_hoursbefore(hours int64, timestamp int64) string {
	return Date_before(3600 * hours, timestamp)
}

//多少分钟后
func Date_minutesafter(minutes int64, timestamp int64) string {
	return Date_after(60 * minutes, timestamp)
}

//多少分钟前
func Date_minutesbefore(minutes int64, timestamp int64) string {
	return Date_before(60 * minutes, timestamp)
}

//多少秒后
func Date_after(seconds int64, timestamp int64) string {
	return time.Unix(timestamp + seconds, 0).In(cstSh).Format(FORMAT_DATE)
}

//多少秒前
func Date_before(seconds int64, timestamp int64) string {
	return time.Unix(timestamp - seconds, 0).In(cstSh).Format(FORMAT_DATE)
}
