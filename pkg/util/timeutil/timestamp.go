package timeutil

import "time"

func GetNowTimeStamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
}
