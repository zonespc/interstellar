package bank

import (
	"fmt"
	"testing"
)

func TestPayedByMonth(t *testing.T) {
	loan := &HouseLoan{
		principal: 112.0 * 10000,
		rate:      5.292,
		n:         360,
	}

	fmt.Printf("=====================\n总贷款  :% 6.f 万\n年利率:% 6.2f %%\n贷款月数:% 6d 月\n---------------------\n", loan.principal/10000, loan.rate, loan.n)
	fmt.Printf("年利率  :% 6.3f %%\n月利率  :% 6.3f %%\n", loan.rate, loan.rate/12)
	refund := PayedByMonth(loan)
	fmt.Printf("---------------------\n每月还款:% .3f 元\n=====================\n", refund)
}

func TestMonthlyInterests(t *testing.T) {
	loan := &HouseLoan{
		principal: 112.0 * 10000,
		rate:      5.292,
		date:      "2019-06-20",
		n:         360,
	}
	realDays := 11
	totalDays := 30
	paid := MonthlyInterests(loan, realDays, totalDays)
	fmt.Println(paid)
}

func TestCalculateMonthlyPayed(t *testing.T) {
	loan := &HouseLoan{
		principal: 112.0 * 10000,
		rate:      5.292,
		date:      "2019-06-20",
		n:         360,
	}
	loanPlan := CalculateMonthlyPayed(loan)
	for i, date := range loanPlan.date {
		fmt.Printf("第%-4d期 %s,% 10.2f,% 10.2f,% 10.2f,% 10.2f\n", i+1, date, loanPlan.left[i], loanPlan.pr[i], loanPlan.in[i], loanPlan.payed[i])
	}
}
