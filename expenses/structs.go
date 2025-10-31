package expenses

import (
	"time"
)

type Expense struct {
	Description string   `json:"description"`
	Amount      int      `json:"amount"`
	Date        string   `json:"date"`
	Category    []string `json:"categories"`
}

func NewExpense(Description string, Amount int, Categories []string) *Expense {
	return &Expense{
		Description: Description,
		Amount:      Amount,
		Date:        Date(time.Now()),
		Category:    Categories,
	}
}

type MonthExpense struct {
	Expenses map[int]Expense `json:"expenses"`
	Budget   int             `json:"budget"`
	Summary  int             `json:"summary"`
}

func NewMonth(amount int) *MonthExpense {
	return &MonthExpense{
		Expenses: make(map[int]Expense, 0),
		Budget:   amount,
		Summary:  0,
	}
}

type ExpenseManager struct {
	MonthExpenses map[time.Month]*MonthExpense
	Total         int
}

func NewExpenseManager() *ExpenseManager {
	k := &ExpenseManager{
		MonthExpenses: make(map[time.Month]*MonthExpense),
		Total:         0,
	}
	_, month, _ := time.Now().Date()
	k.MonthExpenses[month] = NewMonth(1000)
	return k
}
