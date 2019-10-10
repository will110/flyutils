package flyutils

import (
	"strconv"
	"time"
)

/**
描述：获取每月的几号
 */
func GetMonthDay() string {
	dayInt := time.Now().Day()
	return strconv.Itoa(dayInt)
}

/*
描述：获取星期几
返回的格式：
	星期一 = 1
	星期二 = 2
	星期三 = 3
	星期四 = 4
	星期五 = 5
	星期六 = 6
	星期日 = 7
*/
func GetWeekday() string {
	wday := strconv.Itoa(int(time.Now().Weekday()))
	if wday == "0" {
		wday = "7"
	}

	return wday
}