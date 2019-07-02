package calendar

import (
	"fmt"
	"testing"
	"time"
)

func TestPayedMonths(t *testing.T) {
	tm := time.February
	fmt.Println(tm)
	n := 360
	dates := []string{
		//"2019-01-30",
		"2019-06-20",
	}
	for _, date := range dates {
		payedTime := PayedMonths(date, n)
		for i, date := range payedTime.dates {
			fmt.Printf("Date:%s, Days:%d, Total Days:%d\n", date, payedTime.days[i], payedTime.totalDays[i])
		}
	}
}
