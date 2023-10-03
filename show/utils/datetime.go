package utils

import "time"

type DateTime struct {
	tz     string
	format string
}

func (dt *DateTime) SetDateTime(tz string, format string) {
	dt.tz = tz
	dt.format = format
}

func (dt *DateTime) GetDatetime() string {
	time.Local, _ = time.LoadLocation(dt.tz)
	now := time.Now().Format(dt.format)
	return now
}
