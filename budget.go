package go_refactor

import (
	"time"
)

type Budget struct {
	YearMonth string
	Amount    int
}

func NewBudget(yearMonth string, amount int) *Budget {
	return &Budget{YearMonth: yearMonth, Amount: amount}
}

func (budget *Budget) DayBudget() int {
	t, err := time.Parse("200601", budget.YearMonth)
	if err != nil {
		panic(err)
	}
	daysInMonth := time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.Local).Day()
	return budget.Amount / daysInMonth
}

func (budget *Budget) Month() int {
	t, err := time.Parse("200601", budget.YearMonth)
	if err != nil {
		panic(err)
	}
	return int(t.Month())
}
