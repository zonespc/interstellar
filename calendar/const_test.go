package calendar

import (
	"fmt"
	"testing"
	"time"
)

func TestDaysOfMonth(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, v := range m {
		days := DaysOfMonth(time.Month(v), 2016)
		fmt.Println(days)
	}
}
