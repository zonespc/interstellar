package calendar

import "time"

//Layout 时间格式
const (
	Layout = "2006-01-02"
)

//MonthDays 月份天数映射
var MonthDays = map[time.Month]int{
	time.January:   31,
	time.February:  28,
	time.March:     31,
	time.April:     30,
	time.May:       31,
	time.June:      30,
	time.July:      31,
	time.August:    31,
	time.September: 30,
	time.October:   31,
	time.November:  30,
	time.December:  31,
}

//DaysOfMonth 获得某月的天数
func DaysOfMonth(month time.Month, year int) int {
	if month == time.February && IsLeap(year) {
		return 29
	}
	return MonthDays[month]
}

//IsLeap  判断是否为闰年
func IsLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
