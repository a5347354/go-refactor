package go_refactor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBudgetService(t *testing.T) {
	budgetService := BudgetService{}

	t.Run("should return 0 when invalid date", func(t *testing.T) {
		result := budgetService.Query("20230302", "20230301")
		assert.Equal(t, 0, result)
	})

	t.Run("should return 0 when budget data is empty", func(t *testing.T) {
		result := budgetService.Query("20230401", "20230402")
		assert.Equal(t, 0, result)
	})

	t.Run("query single month", func(t *testing.T) {
		result := budgetService.Query("20230101", "20230131")
		assert.Equal(t, 310, result)
	})

	t.Run("cross 2 months", func(t *testing.T) {
		result := budgetService.Query("20230131", "20230202")
		assert.Equal(t, 50, result)
	})

	t.Run("cross 3 months", func(t *testing.T) {
		result := budgetService.Query("20230131", "20230302")
		assert.Equal(t, 560+10+60, result)
	})
}
