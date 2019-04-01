package funTime

import (
	"time"
)

func Time() int64 {
	return time.Now().Unix()
}