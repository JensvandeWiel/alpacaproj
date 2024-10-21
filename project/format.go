package project

import "time"

var TimestampFormat = "20060102150405"

func GenerateTimestamp() string {
	return time.Now().UTC().Format(TimestampFormat)
}
