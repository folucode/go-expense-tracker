package expense

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var db = "expenses.json"

var months = []string{
	"January",
	"Febuary",
	"March", "April", "May", "June", "July", "August", "September", "October", "November", "December",
}

type Expense struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Month       string  `json:"month"`
	Year        int     `json:"year"`
}

func LoadExpenses() []Expense {
	file, err := os.ReadFile(db)
	if err != nil {
		return []Expense{}
	}
	var expenses []Expense
	json.Unmarshal(file, &expenses)
	return expenses
}

func AddExpense(expenses *[]Expense, title, desc string, amount float64) {
	id := 1
	if len(*expenses) > 0 {
		id = (*expenses)[len(*expenses)-1].ID + 1
	}
	newExpense := Expense{
		ID: id, Title: title, Description: desc, Amount: amount,
		Month: months[int(time.Now().Month())-1], Year: time.Now().Year(),
	}
	*expenses = append(*expenses, newExpense)
	fmt.Println("Expense added.")
}

func UpdateExpense(expense *Expense, updateData []string) {
	for _, fieldVal := range updateData {
		parts := strings.SplitN(fieldVal, ":", 2)
		if len(parts) != 2 {
			fmt.Printf("Warning: Skipping invalid update format '%s'. Use 'field:value'.\n", fieldVal)
			continue
		}
		field := parts[0]
		value := parts[1]
		switch field {
		case "title":
			expense.Title = value
		case "description":
			expense.Description = value
		case "amount":
			amount, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Printf("Error: Invalid amount format for '%s'.\n", value)
				continue
			}
			expense.Amount = amount
		case "month":
			expense.Month = value
		case "year":
			year, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				fmt.Printf("Error: Invalid year format for '%s'.\n", value)
				continue
			}
			expense.Year = int(year)
		default:
			fmt.Printf("Warning: Unknown field '%s'.\n", field)
		}
	}
	fmt.Printf("Final expense: %+v\n", *expense)
}

func FindExpenseIndex(id int, expenses *[]Expense) (int, bool) {
	for i, e := range *expenses {
		if e.ID == id {
			return i, true
		}
	}
	return -1, false
}

func DeleteExpense(expenses *[]Expense, expenseToDelete Expense) {
	updatedSlice := slices.DeleteFunc(*expenses, func(e Expense) bool {
		return e.ID == expenseToDelete.ID
	})
	*expenses = updatedSlice
}

func ViewExpenses(expenses *[]Expense) {
	fmt.Println("id, title, description, amount, month, year")
	for _, e := range *expenses {
		fmt.Printf("%d, %s, %s, %.2f, %s, %d\n", e.ID, e.Title, e.Description, e.Amount, e.Month, e.Year)
	}
}

func SummarizeExpenses(expenses *[]Expense, month string) {
	totalExpenses := 0.00
	isMonthProvided := len(month) > 0
	for _, e := range *expenses {
		if !isMonthProvided || strings.EqualFold(month, e.Month) {
			totalExpenses += e.Amount
		}
	}
	fmt.Printf("Total expenses: $%.2f\n", totalExpenses)
}

func SaveExpenses(expenses []Expense) {
	data, _ := json.MarshalIndent(expenses, "", "  ")
	os.WriteFile(db, data, 0644)
}
