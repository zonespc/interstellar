package bank

import (
	"fmt"
	"time"
)

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

//PayedTime 贷款时间相关结构
type PayedTime struct {
	dates     []string
	days      []int
	totalDays []int
}

//PayedMonths year——年 month——月 day——日 n——总月数
func PayedMonths(date string, n int) (payedTime *PayedTime) {
	payedTime = &PayedTime{
		dates:     []string{},
		days:      []int{},
		totalDays: []int{},
	}
	tdate, err := time.Parse(Layout, date)
	if err != nil {
		fmt.Println("parse err=", err)
	}
	y := tdate.Year()
	m := tdate.Month()
	d := tdate.Day()
	Days := make([]int, n)
	nDays := DaysOfMonth(m, y) - d + 1
	totalDays := DaysOfMonth(m, y)
	Days[0] = nDays
	t := tdate.AddDate(0, 0, -d+1)
	for i := 1; i <= n; i++ {
		t := t.AddDate(0, i, 0)
		payedTime.dates = append(payedTime.dates, t.Format(Layout))
		payedTime.days = append(payedTime.days, nDays)
		payedTime.totalDays = append(payedTime.totalDays, totalDays)
		nDays = DaysOfMonth(t.Month(), t.Year())
		totalDays = nDays
	}

	return payedTime
}
