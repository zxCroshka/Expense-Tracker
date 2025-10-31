package expenses

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
	"time"
)

func (m *ExpenseManager) ConfigBudget(amount int) {

	if err := CreateIfNotExists("expenses.json"); err != nil {
		log.Fatal(err)
	}
	if err := JSONtoStruct(&m); err != nil {
		log.Fatal()
	}
	month := GetMonth()
	if v, ok := m.MonthExpenses[month]; ok {
		v.Budget = amount

	} else {
		m.MonthExpenses[month] = NewMonth(amount)
	}
	if err := StructToJSON(m); err != nil {
		log.Fatal(err)
	}
}

func (m *ExpenseManager) Budget(month time.Month) {
	if err := JSONtoStruct(&m); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Budget for %s: %d\n", month.String(), m.MonthExpenses[month].Budget)
}

func (m *ExpenseManager) Add(Description string, Amount int, Categories []string) error {
	if err := CreateIfNotExists("expenses.json"); err != nil {
		log.Fatal(err)
	}
	month := GetMonth()
	if err := JSONtoStruct(&m); err != nil {
		log.Fatal()
	}
	l := len(m.MonthExpenses[month].Expenses)
	for k := range m.MonthExpenses[month].Expenses {
		l = max(l, k)
	}

	if Amount <= m.MonthExpenses[month].Budget {
		m.MonthExpenses[month].Expenses[l+1] = *NewExpense(Description, Amount, Categories)
		m.MonthExpenses[month].Budget -= Amount
		m.MonthExpenses[month].Summary += Amount
		m.Total += Amount
	} else {
		log.Println("not enough money")
		return ErrNOM
	}
	if err := StructToJSON(m); err != nil {
		log.Fatal(err)
	}
	log.Printf("Task added successfully (ID: %d)", l+1)
	return nil
}

func (m *ExpenseManager) Delete(id int) error {
	month := GetMonth()
	if err := JSONtoStruct(&m); err != nil {
		log.Fatal(err)
	}
	v, ok := m.MonthExpenses[month].Expenses[id]
	if !ok {
		log.Println("There isn't expense with this id at current Month")
		return ErrNotExist
	}
	m.Total -= v.Amount
	m.MonthExpenses[month].Budget += v.Amount
	m.MonthExpenses[month].Summary -= v.Amount
	delete(m.MonthExpenses[month].Expenses, id)

	if err := StructToJSON(m); err != nil {
		log.Fatal(err)
	}
	log.Printf("Task deleted (ID: %d)", id)
	return nil
}

func (m *ExpenseManager) GetTotal() int {
	if err := JSONtoStruct(&m); err != nil {
		log.Fatal(err)
	}
	return m.Total
}

func (m *ExpenseManager) GetMonthSummary(month time.Month) (int, error) {
	if err := JSONtoStruct(&m); err != nil {
		log.Fatal(err)
	}
	v, ok := m.MonthExpenses[month]
	if !ok {
		log.Println("No Expenses at this Month")
		return 0, ErrMonth
	}
	return v.Summary, nil
}

func (m *ExpenseManager) ListMonth(month time.Month, categories []string) {
	if err := JSONtoStruct(&m); err != nil {
		log.Fatal(err)
	}
	keys := make([]int, 0)
	for k := range m.MonthExpenses[month].Expenses {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	titles := []string{"ID", "Date", "Description", "Amount", "Categories"}
	titleLens := make([]int, len(titles))
	for i, v := range titles {
		titleLens[i] = len(v)
	}
	maxLen := make([]int, len(titleLens))
	copy(maxLen, titleLens)
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		temp := m.MonthExpenses[month].ToLen(k)
		for j := 0; j < len(temp); j++ {
			maxLen[j] = max(temp[j], maxLen[j])
		}
	}
	summary, err := m.GetMonthSummary(month)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Month: %s,	Budget: %d, Summary of Expenses: %v\n", month.String(), m.MonthExpenses[month].Budget, summary)
	fmt.Printf("%v%v  %v%v  %v%v  %v%v  %v%v\n",
		titles[0], strings.Repeat(" ", maxLen[0]-titleLens[0]),
		titles[1], strings.Repeat(" ", maxLen[1]-titleLens[1]),
		titles[2], strings.Repeat(" ", maxLen[2]-titleLens[2]),
		titles[3], strings.Repeat(" ", maxLen[3]-titleLens[3]),
		titles[4], strings.Repeat(" ", maxLen[4]-titleLens[4]),
	)
	temp := make([]int, 0)
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		e := m.MonthExpenses[month].Expenses[k]
		if len(categories) != 0 {
			flag := true
			for _, v := range categories {
				if !slices.Contains(e.Category, v) || len(e.Category) == 0 {
					flag = false
					break
				}
			}
			if flag {
				temp = append(temp, k)
			}
		} else {
			temp = keys
		}
	}

	for i := 0; i < len(temp); i++ {
		k := temp[i]
		e := m.MonthExpenses[month].Expenses[k]
		lens := m.MonthExpenses[month].ToLen(k)
		fmt.Printf("%v%v  %v%v  %v%v  %v%v  %v%v\n",
			k, strings.Repeat(" ", maxLen[0]-lens[0]),
			e.Date, strings.Repeat(" ", maxLen[1]-lens[1]),
			e.Description, strings.Repeat(" ", maxLen[2]-lens[2]),
			e.Amount, strings.Repeat(" ", maxLen[3]-lens[3]),
			strings.Join(e.Category, "; "), strings.Repeat(" ", maxLen[4]-lens[4]),
		)
	}
}

func (m *ExpenseManager) List(categories []string) {
	for month := range m.MonthExpenses {
		m.ListMonth(month, categories)
	}

}

func (m *ExpenseManager) CSVExport() error {

	if err := CreateIfNotExists("expenses.csv"); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("expenses.csv", os.O_RDWR, 0777)
	defer func(){
		if err := file.Close(); err != nil{
			log.Fatal(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	if err := JSONtoStruct(&m); err != nil {
		log.Fatal(err)
	}

	titles := []string{"ID", "Date", "Description", "Amount", "Categories"}
	if err := writer.Write(titles); err != nil {
		log.Fatal(err)
	}
	for month := range m.MonthExpenses {
		keys := make([]int, 0)
		for k := range m.MonthExpenses[month].Expenses {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for i := 0; i < len(keys); i++ {
			k := keys[i]
			e := m.MonthExpenses[month].Expenses[k]
			curs := fmt.Sprintf("%v %v %v %v %v",
				k,
				e.Date,
				e.Description,
				e.Amount,
				strings.Join(e.Category, ";"),
			)
			tmp := strings.Split(curs, " ")
			if err := writer.Write(tmp); err != nil{
				log.Fatal(err)
			}
		}
		
	}

	log.Println("CSV file exported successfully, name of file: 'expenses.csv'")
	return nil
}
