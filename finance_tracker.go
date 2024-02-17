// hands-on experience with Go's syntax, control structures, data types, slices, maps and structs

package main

import (
	"fmt"
	"strconv"
)

type IncomeType string

const (
	income  IncomeType = "income"
	expense IncomeType = "expense"
)

// TODO: Define a struct for FinanceRecord with fields for amount, description, category, and type (income or expense)
type FinanceRecord struct {
	amount      float64
	description string
	category    string
	income_type IncomeType
}

// TODO: Define global slice to hold finance records
var financeRecords []FinanceRecord

// TODO: Implement function to add a finance record
func addFinanceRecord(amt float64, desc string, category string, inc IncomeType) {
	// Implement function
	temp_record := FinanceRecord{
		amount:      amt,
		description: desc,
		category:    category,
		income_type: inc,
	}
	financeRecords = append(financeRecords, temp_record)
}

// TODO: Implement function to print summary (total income, total expenses, net balance)
func printSummary() {
	// Implement function
	fmt.Println("The summary is as follows:-")
	var final_income float64
	var total_expenses float64
	var net_balance float64
	for _, v := range financeRecords {
		if v.income_type == "income" {
			final_income += v.amount
			net_balance += v.amount
		} else if v.income_type == "expense" {
			total_expenses += v.amount
			final_income -= v.amount
		}
	}
	fmt.Printf("Total Income:- %f\n", final_income)
	fmt.Printf("Total Expenses:- %f\n", total_expenses)
	fmt.Printf("NetBalance:- %f\n", net_balance)
}

// TODO: Implement function to print records filtered by type (income/expense)
func printRecordsByType(income_type string) {
	// Implement function
	for _, v := range financeRecords {
		if string(v.income_type) == income_type {
			fmt.Printf("Amount:-%f Description:-%s Category:-%s incomeType:-%v \n", v.amount, v.description, v.category, v.income_type)
		}
	}
}

// TODO: Implement function for basic financial calculations (e.g., estimate next month savings)
func calculateSavings() {
	// Implement function
	var result float64 = 0
	for _, v := range financeRecords {
		if v.income_type == income {
			result += v.amount
		} else {
			result -= v.amount
		}
	}
	fmt.Printf("The savings are %f\n", result)
}

// TODO: Implement function to get user input and return as appropriate data type
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

// TODO: Implement main logic for user interaction
func main() {
	for {
		fmt.Println("\nPersonal Finance Tracker")
		fmt.Println("1. Add Record")
		fmt.Println("2. View Summary")
		fmt.Println("3. View Records by Type")
		fmt.Println("4. Calculate Savings")
		fmt.Println("5. Exit")
		choice := getUserInput("Choose an option: ")

		switch choice {
		case "1":
			// TODO: Implement adding a record
			var amount float64
			var description string
			var category string
			var income_type IncomeType
			fmt.Println("please enter the amount")
			for {
				var temp_input string
				fmt.Scanf("%s", &temp_input)
				if _, err := strconv.ParseFloat(temp_input, 64); err != nil {
					fmt.Println("Invalid operation, please enter a valid amount")
				} else {
					amount, _ = strconv.ParseFloat(temp_input, 64)
					break
				}
			}
			fmt.Println("please enter the description")
			fmt.Scanf("%s", &description)
			fmt.Println("please enter the category")
			fmt.Scanf("%s", &category)
			fmt.Println("please enter the income type, please note that it can only be expense or income")
			for {
				var temp_input string
				fmt.Scanf("%s", &temp_input)
				if temp_input != string(income) && temp_input != string(expense) {
					fmt.Println("Invalid operation, please enter a valid amount")
				} else {
					if temp_input == string(income) {
						income_type = income
					} else if temp_input == string(expense) {
						income_type = expense
					}
					break
				}
			}
			addFinanceRecord(amount, description, category, income_type)
		case "2":
			// TODO: Implement viewing summary
			printSummary()
		case "3":
			// TODO: Implement viewing records by type
			var income_type IncomeType
			for {
				var temp_input string
				fmt.Scanf("%s", &temp_input)
				if temp_input != string(income) && temp_input != string(expense) {
					fmt.Println("Invalid operation, please enter a valid amount")
				} else {
					if temp_input == string(income) {
						income_type = income
					} else if temp_input == string(expense) {
						income_type = expense
					}
					break
				}
			}
			printRecordsByType(string(income_type))
		case "4":
			// TODO: Implement calculating savings
			calculateSavings()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
