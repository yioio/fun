package funTime

import (
	"time"
	// "log"
)

func init() {
	cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
}

//dateTimeFormat 格式如："2006-01-02 15:04:05" 
func Datetime(timestamp int64) string {
	return time.Unix(timestamp, 0).In(cstSh).Format(FORMAT_DATETIME)
}

//当前日期时间
func Datetime_now() string {
	return Datetime_today()
}

//当前日期时间
func Datetime_today() string {
	return time.Now().In(cstSh).Format(FORMAT_DATETIME)
}

//明天时间
func Datetime_tomorrow() string {
	
	return time.Unix(Time() + 86400, 0).In(cstSh).Format(FORMAT_DATETIME)
}

//多少天后
func Datetime_daysafter(days int64, timestamp int64) string {
	return Datetime_after(86400 * days, timestamp)
}

//多少天前
func Datetime_daysbefore(days int64, timestamp int64) string {
	return Datetime_before(86400 * days, timestamp)
}

//多少小时后
func Datetime_hoursafter(hours int64, timestamp int64) string {
	return Datetime_after(3600 * hours, timestamp)
}

//多少小时前
func Datetime_hoursbefore(hours int64, timestamp int64) string {
	return Datetime_before(3600 * hours, timestamp)
}

//多少分钟后
func Datetime_minutesafter(minutes int64, timestamp int64) string {
	return Datetime_after(60 * minutes, timestamp)
}

//多少分钟前
func Datetime_minutesbefore(minutes int64, timestamp int64) string {
	return Datetime_before(60 * minutes, timestamp)
}

//多少秒后
func Datetime_after(seconds int64, timestamp int64) string {
	return time.Unix(timestamp + seconds, 0).In(cstSh).Format(FORMAT_DATETIME)
}

//多少秒前
func Datetime_before(seconds int64, timestamp int64) string {
	return time.Unix(timestamp - seconds, 0).In(cstSh).Format(FORMAT_DATETIME)
}
