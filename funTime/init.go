package funTime

import (
	"time"
)

var (
	cstSh *time.Location
)

var FORMAT_DATE = "2006-01-02"
var FORMAT_DATETIME string = "2006-01-02 15:04:05"

/**
	PHP的date("Y-m-d", timestamp) == funTime.Date(timestamp)			
	PHP的date("Y-m-d H:i:s", timestamp) == funTime.Datetime(timestamp)	
	time() == funTime.Time()
*/