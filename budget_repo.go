package go_refactor

type BudgetRepo struct{}

func (repo *BudgetRepo) GetAll() []*Budget {
	return []*Budget{
		NewBudget("202301", 310),
		NewBudget("202302", 560),
		NewBudget("202303", 930),
	}
}
