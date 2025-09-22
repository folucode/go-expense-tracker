package main

import (
	"fmt"
	"os"
	"strconv"

	"go-expense-tracker/internal/expense"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: expense [add|update|delete|summary|list]")
		return
	}

	command := os.Args[1]
	expenses := expense.LoadExpenses()

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: expense add <title> <description> <amount>")
			return
		}
		title := os.Args[2]
		description := os.Args[3]
		amount, _ := strconv.ParseFloat(os.Args[4], 64)
		expense.AddExpense(&expenses, title, description, amount)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: expense update <id> <field>:<value> [<field>:<value>...]")
			return
		}

		id, err := strconv.ParseInt(os.Args[2], 10, 0)
		if err != nil {
			fmt.Println("Error: Invalid ID format.")
			return
		}

		updateData := os.Args[3:]

		expenseIndex, found := expense.FindExpenseIndex(int(id), &expenses)
		if !found {
			fmt.Println("Error: Expense not found.")
			return
		}

		expense.UpdateExpense(&expenses[expenseIndex], updateData)
		fmt.Printf("Expense with ID %d updated successfully.\n", int(id))
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: expense delete <id>")
			return
		}

		id, err := strconv.ParseInt(os.Args[2], 10, 0)
		if err != nil {
			fmt.Println("Error: Invalid ID format.")
			return
		}

		expenseIndex, found := expense.FindExpenseIndex(int(id), &expenses)
		if !found {
			fmt.Println("Error: Expense not found.")
			return
		}
		expense.DeleteExpense(&expenses, expenses[expenseIndex])
		fmt.Printf("Expense with ID %d deleted successfully.\n", int(id))
	case "list":
		if len(os.Args) < 2 {
			fmt.Println("Usage: expense list")
			return
		}
		expense.ViewExpenses(&expenses)
	case "summary":
		var month string
		if len(os.Args) > 2 {
			month = os.Args[2]
		}
		expense.SummarizeExpenses(&expenses, month)
	default:
		fmt.Println("Unknown command:", command)
	}
	expense.SaveExpenses(expenses)
}
