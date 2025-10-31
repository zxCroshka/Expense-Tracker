package expenses

import (
	"strings"
	"time"
)

func Date(t time.Time) string {
	return t.Format(time.DateOnly)
}

func GetMonth() time.Month {
	_, month, _ := time.Now().Date()
	return month
	
}

func (m *MonthExpense) ToLen(id int) []int {
	column := make([]int, 0)
	lID := 0
	temp := id
	for temp != 0 {
		temp = temp / 10
		lID++
	}
	column = append(column, lID)
	column = append(column, len(m.Expenses[id].Date))
	column = append(column, len(m.Expenses[id].Description))
	lAmount := 0
	temp = m.Expenses[id].Amount
	for temp != 0 {
		temp = temp / 10
		lAmount++
	}
	column = append(column, lAmount)
	k := strings.Join(m.Expenses[id].Category,"; ")
	column = append(column, len(k))
	return column
}
