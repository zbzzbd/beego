package utils

import "time"

var DefaultTimeLayout = "2006-01-02 15:04"
var YMDHIS = "2006-01-02 15:04:05"

func FormatTime(src string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation("2006-01-02 15:04", src, loc)
}
