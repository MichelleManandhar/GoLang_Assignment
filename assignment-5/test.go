// Create a banking application. Your bank should have your customer data.
// The bank should be able to provide information about its customers like total number of customers and print all customers data.
// Then your customer should have the capability to withdraw money, deposit money, print balance and close the account.
// Also your code should be able to handle adding new accounts
// package main

// type balance struct {
// 	amount int
// }

// type customer struct {
// 	FirstName string
// 	LastName  string
// 	age       int
// 	phone     int
// 	bal       balance
// }

// func (bal balance) depositBalance() {

// }

// func (bal balance) withdrawBalance() {

// }

// func main() {
// 	cus1 := customer{
// 		"Michelle",
// 		"Manandhar",
// 		25,
// 		9282257777,
// 		balance{
// 			50000,
// 		},
// 	}
// 	cus1.bal.depositBalance()
// 	cus1.bal.withdrawBalance()

// }
//..................................................................

//......................................................................

//package main

import (
	"fmt"
)

type Customer struct {
	Name    string
	Account Account
}

type Account struct {
	AccountNumber int
	Balance       float64
}

type Bank struct {
	Customers []Customer
}

func (b *Bank) AddCustomer(name string, accountNumber int, initialBalance float64) {
	customer := Customer{
		Name: name,
		Account: Account{
			AccountNumber: accountNumber,
			Balance:       initialBalance,
		},
	}
	b.Customers = append(b.Customers, customer)
}

func (b *Bank) GetTotalCustomers() int {
	return len(b.Customers)
}

func (b *Bank) PrintAllCustomers() {
	for _, customer := range b.Customers {
		fmt.Printf("Customer Name: %s\n", customer.Name)
		fmt.Printf("Account Number: %d\n", customer.Account.AccountNumber)
		fmt.Printf("Account Balance: %.2f\n\n", customer.Account.Balance)
	}
}

func (c *Customer) Deposit(amount float64) {
	c.Account.Balance += amount
	fmt.Printf("Deposit successful. New balance: %.2f\n", c.Account.Balance)
}

func (c *Customer) Withdraw(amount float64) {
	if amount <= c.Account.Balance {
		c.Account.Balance -= amount
		fmt.Printf("Withdrawal successful. New balance: %.2f\n", c.Account.Balance)
	} else {
		fmt.Println("Insufficient funds.")
	}
}

func (c *Customer) PrintBalance() {
	fmt.Printf("Current balance: %.2f\n", c.Account.Balance)
}

func (b *Bank) CloseAccount(accountNumber int) {
	for i, customer := range b.Customers {
		if customer.Account.AccountNumber == accountNumber {
			// Remove the customer from the slice
			b.Customers = append(b.Customers[:i], b.Customers[i+1:]...)
			fmt.Println("Account closed successfully.")
			return
		}
	}
	fmt.Println("Account not found.")
}

func main() {
	bank := Bank{}

	for {
		fmt.Println("\n************ Bank Application ************")
		fmt.Println("1. Add Customer")
		fmt.Println("2. Total Number of Customers")
		fmt.Println("3. Print All Customer Data")
		fmt.Println("4. Select Customer")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter customer name: ")
			var name string
			fmt.Scanln(&name)

			fmt.Print("Enter account number: ")
			var accountNumber int
			fmt.Scanln(&accountNumber)

			fmt.Print("Enter initial balance: ")
			var initialBalance float64
			fmt.Scanln(&initialBalance)

			bank.AddCustomer(name, accountNumber, initialBalance)
			fmt.Println("Customer added successfully.")

		case 2:
			totalCustomers := bank.GetTotalCustomers()
			fmt.Printf("Total number of customers: %d\n", totalCustomers)

		case 3:
			fmt.Println("All Customers:")
			bank.PrintAllCustomers()

		case 4:
			fmt.Print("Enter account number: ")
			var accountNumber int
			fmt.Scanln(&accountNumber)

			selectedCustomer := findCustomerByAccountNumber(bank.Customers, accountNumber)
			if selectedCustomer == nil {
				fmt.Println("Customer not found.")
				continue
			}
			for {
				fmt.Println("\n************ Customer Operations ************")
				fmt.Println("1. Deposit")
				fmt.Println("2. Withdraw")
				fmt.Println("3. Print Balance")
				fmt.Println("4. Close Account")
				fmt.Println("5. Go back to Bank Menu")
				fmt.Print("Enter your choice: ")

				var customerChoice int
				fmt.Scanln(&customerChoice)

				switch customerChoice {
				case 1:
					fmt.Print("Enter amount to deposit: ")
					var amount float64
					fmt.Scanln(&amount)
					selectedCustomer.Deposit(amount)
				case 2:
					fmt.Print("Enter amount to withdraw: ")
					var amount float64
					fmt.Scanln(&amount)
					selectedCustomer.Withdraw(amount)

				case 3:
					selectedCustomer.PrintBalance()

				case 4:
					bank.CloseAccount(accountNumber)
					break

				case 5:
					break

				default:
					fmt.Println("Invalid choice.")
				}

				if customerChoice == 5 {
					break
				}
			}
		case 5:
			fmt.Println("Thank you for using the banking application.")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}


selectedCustomer := findCustomerByAccountNumber(bank.Customers, accountNumber)
			if selectedCustomer == nil {
				fmt.Println("Customer not found.")
				continue
			}

func findCustomerByAccountNumber(customers []Customer, accountNumber int) *Customer {
	for i := range customers {
		if customers[i].Account.AccountNumber == accountNumber {
			return &customers[i]
		}
	}
	return nil
}
