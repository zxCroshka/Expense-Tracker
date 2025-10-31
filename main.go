package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/zxCroshka/Expense-Tracker/expenses"
)

func main() {
	manager := expenses.NewExpenseManager()
	cmd := os.Args[1]
	switch cmd {
	case "add":
		description := os.Args[3]
		amount, err := strconv.Atoi(os.Args[5])
		if err != nil {
			log.Fatal("Invalid amount")
		}
		categories := os.Args[7:]
		if err := manager.Add(description, amount, categories); err != nil {
			if errors.Is(err, expenses.ErrNOM) {
				log.Println("Failed to add expense")
			} else {
				log.Fatal()
			}
		}
	case "list":
		if len(os.Args) == 2 {
			manager.List()
		} else {
			month,err := strconv.Atoi(os.Args[4])
			if err != nil{
				log.Fatal("Invalid month")
			}
			
			manager.ListMonth(time.Month(month))

		}
	case "summary":
		if len(os.Args) == 2{
			fmt.Printf("Total expenses: %d\n",manager.GetTotal())
		}else{
			month,err := strconv.Atoi(os.Args[3])
			if err != nil{
				log.Fatal("Invalid month")
			}
			summary, err := manager.GetMonthSummary(time.Month(month));
			if err != nil{
				log.Fatal(err)
			}
			fmt.Printf("Total Expenses for %s: %d",time.Month(month).String(),summary)
		}
	case "delete":
		id, err := strconv.Atoi(os.Args[3])
		if err != nil{
			log.Fatal("Invalid id")
		}
		manager.Delete(id)
	case "config-budget":
		amount,err := strconv.Atoi(os.Args[3])
		if err != nil{
			log.Fatal("Invalid amount")
		}
		manager.ConfigBudget(amount)
	}

}
