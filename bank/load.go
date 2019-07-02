package bank

import (
	"math"
)

//HouseLoan 房贷数据结构
// —principal 贷款总额
// -rate 年利率
// -date 贷款时间，"2019-06-20"
// -n 贷款月数，360
type HouseLoan struct {
	principal float64
	rate      float64
	date      string
	n         int
}

//PayedByMonth 等额本息计算
// <-输入：房贷数据结构
// ->输出：每月还款额度
func PayedByMonth(loan *HouseLoan) float64 {
	monthlyRate := loan.rate / 12 * 0.01
	p := math.Pow(1+monthlyRate, float64(loan.n))
	return loan.principal * monthlyRate * p / (p - 1)
}

//MonthlyInterests 每月付款利息
// -loan: 房贷数据结构
// -realDays: 上月贷款天数
// -totalDays: 上月总天数
func MonthlyInterests(loan *HouseLoan, realDays int, totalDays int) float64 {
	monthlyRate := loan.rate / 12 * 0.01
	in := loan.principal * monthlyRate / float64(totalDays) * float64(realDays)
	//fmt.Printf("%.2f %d %d %.2f %f\n", loan.principal, realDays, totalDays, in, monthlyRate)
	return in
}

//LoanPlan 贷款计划
// -date  还款时间
// -left  上月还款后剩余本金
// -pr    上月还款本金
// -in    上月还款利息
// -payed 上月还款本息
type LoanPlan struct {
	date  []string
	left  []float64
	pr    []float64
	in    []float64
	payed []float64
}

//CalculateMonthlyPayed 计算每月还款额
func CalculateMonthlyPayed(loan *HouseLoan) *LoanPlan {
	loanPlan := &LoanPlan{
		date:  []string{},
		left:  []float64{},
		pr:    []float64{},
		in:    []float64{},
		payed: []float64{},
	}
	refund := math.Round(PayedByMonth(loan)*100) / 100
	payedTime := PayedMonths(loan.date, loan.n)
	for i, date := range payedTime.dates {
		mInterests := MonthlyInterests(loan, payedTime.days[i], payedTime.totalDays[i])
		In := math.Round(mInterests*100) / 100
		Pr := refund - In
		payed := Pr + In
		if In <= 0 {
			Pr = 0
			payed = 0
		}
		if In+loan.principal < refund {
			payed = In + loan.principal
			Pr = loan.principal
		}
		loan.principal -= Pr
		if loan.principal < 0 {
			loan.principal = 0
		}
		loanPlan.date = append(loanPlan.date, date)
		loanPlan.left = append(loanPlan.left, loan.principal)
		loanPlan.pr = append(loanPlan.pr, Pr)
		loanPlan.in = append(loanPlan.in, In)
		loanPlan.payed = append(loanPlan.payed, payed)
	}
	return loanPlan
}
