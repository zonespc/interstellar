package calendar

import (
	"fmt"
	"time"
)

//PayedTime 贷款时间相关结构
type PayedTime struct {
	dates     []string
	days      []int
	totalDays []int
}

//PayedMonths year——年 month——月 day——日 n——总月数
func PayedMonths(date string, nMonths int) (payedTime *PayedTime) {
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
	Days := make([]int, nMonths)
	nDays := DaysOfMonth(m, y) - d + 1
	totalDays := DaysOfMonth(m, y)
	Days[0] = nDays
	t := tdate.AddDate(0, 0, -d+1)
	for i := 1; i < nMonths; i++ {
		t := t.AddDate(0, i, 0)
		payedTime.dates = append(payedTime.dates, t.Format(Layout))
		payedTime.days = append(payedTime.days, nDays)
		payedTime.totalDays = append(payedTime.totalDays, totalDays)
		nDays = DaysOfMonth(t.Month(), t.Year())
		totalDays = nDays
	}

	return payedTime
}
