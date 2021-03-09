package util

import "time"

func Gen() (ID int64) {
	return time.Now().UnixNano()
}
