package main

import "fmt"

type Bank struct {
	//field called Customers which has slie of ustomer object
	Cust []Customer
}

type Customer struct {
	//initializing info in struct
	Name    string
	Account int
	Balance float64
}

func (add *Bank) AddCustomer(customerName string, accountNumber int, startingBalance float64) {
	//* as value is shared to store in slice Customer
	addCustomer := Customer{
		Name:    customerName,
		Account: accountNumber,
		Balance: startingBalance,
	}
	add.Cust = append(add.Cust, addCustomer)
	fmt.Println("Customer added successfully.")
	fmt.Println("The total number of customer is", len(add.Cust), ".")
	fmt.Println("----------------------------------------------------------------------------")
}

func (view *Bank) ViewCustomers() {
	for _, value := range view.Cust {
		//loop through customer to print info
		fmt.Printf("Customer : %s\n", value.Name)
		fmt.Printf("Account number: %d\n", value.Account)
		fmt.Printf("Account Balance: %.2f\n", value.Balance)
		fmt.Println("-----------------------------------")
	}
	fmt.Println("----------------------------------------------------------------------------")

}

func main() {
	//create instance of bank
	bank := Bank{}
	for {
		//for to iterate to use the following options
		fmt.Println("Welcome to Banking Application")
		fmt.Println("1. Add Customer")
		fmt.Println("2. View All Customer's Data")
		fmt.Println("3. Make Changes To Account")
		fmt.Println("4. Exit")
		fmt.Println("Enter your choice:")

		//take input from user
		var choice int
		fmt.Scanln(&choice)
		var customerName string
		var accountNumber int
		var startingBalance float64

		//use switch stmt
		switch choice {
		case 1:
			//Adding customer
			fmt.Println("Enter customer name:")
			fmt.Scan(&customerName)

			fmt.Println("Enter account number:")
			fmt.Scan(&accountNumber)

			fmt.Println("Enter customer's account balance")
			fmt.Scanln(&startingBalance)

			//pass the input to the function
			bank.AddCustomer(customerName, accountNumber, startingBalance)

		case 2:
			fmt.Println("The list of all customers:")
			bank.ViewCustomers()

		case 3:
			fmt.Println("Enter the bank account of customer to make changes:")

		case 4:
			fmt.Println("You are exisiting from Banking Application.")
			fmt.Println("----------------------------------------------------------------------------")
			return

		}
	}

}
