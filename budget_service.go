package go_refactor

import (
	"math"
	"time"
)

type BudgetService struct{}

func NewBudgetService() *BudgetService {
	return &BudgetService{}
}

func (s BudgetService) Query(start, end string) int {
	startDate, err := time.Parse("20060102", start)
	if err != nil {
		return 0
	}
	endDate, err := time.Parse("20060102", end)
	if err != nil {
		return 0
	}
	if endDate.Before(startDate) {
		return 0
	}
	budgetRepo := BudgetRepo{}
	budgets := budgetRepo.GetAll()
	keys := s.getKeys(startDate, endDate)
	var searchResult []*Budget
	for _, budget := range budgets {
		for _, key := range keys {
			if key == budget.YearMonth {
				searchResult = append(searchResult, budget)
			}
		}
	}
	if len(searchResult) == 0 {
		return 0
	}
	if len(keys) == 1 {
		days := int(math.Ceil(endDate.Sub(startDate).Hours()/24) + 1)
		return searchResult[0].DayBudget() * days
	}
	var totalBudget int
	for i, budget := range searchResult {
		if i == 0 {
			t, err := time.Parse("200601", budget.YearMonth)
			if err != nil {
				return 0
			}
			monthEnd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 1, -1)
			days := int(math.Ceil(monthEnd.Sub(startDate).Hours()/24) + 1)
			totalBudget += days * budget.DayBudget()
		} else if i == len(searchResult)-1 {
			t, err := time.Parse("200601", budget.YearMonth)
			if err != nil {
				return 0
			}
			monthStart := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
			days := int(math.Ceil(endDate.Sub(monthStart).Hours() / 24))
			totalBudget += days * budget.DayBudget()
		} else {
			totalBudget += budget.Amount
		}
	}
	return totalBudget
}

func (s *BudgetService) getKeys(startDate, endDate time.Time) []string {
	startDate = time.Date(startDate.Year(), startDate.Month(), 1, 0, 0, 0, 0, time.UTC)
	endDate = time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, time.UTC)

	var result []string
	for endDate.Unix() >= startDate.Unix() {
		result = append(result, startDate.Format("200601"))
		startDate = startDate.AddDate(0, 1, 0)
	}

	return result
}
