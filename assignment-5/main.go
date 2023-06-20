// Banking Application
package main

import "fmt"

type Bank struct {
	//field called Customers which has slice of ustomer object
	Cust []Customer
}

type Customer struct {
	//initializing info in struct
	Name string
	Acc  Account
	// AccountNum int
	// Balance float64
}

type Account struct {
	AccountNum int
	Balance    float64
}

func (add *Bank) AddCustomer(customerName string, accountNumber int, startingBalance float64) {
	//* as value is shared to store in slice Customer
	addCustomer := Customer{
		Name: customerName,
		Acc: Account{
			AccountNum: accountNumber,
			Balance:    startingBalance,
		},
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
		fmt.Printf("Account number: %d\n", value.Acc.AccountNum)
		fmt.Printf("Account Balance: %.2f\n", value.Acc.Balance)
		fmt.Println("-----------------------------------")
	}
	fmt.Println("----------------------------------------------------------------------------")

}

func FindCustomer(cust []Customer, acnum int) *Customer {
	var temp *Customer
	for i := range cust {
		if cust[i].Acc.AccountNum == acnum {
			fmt.Println("Showing the information of customer", cust[i].Name)
			temp = &cust[i]
		}
	}
	fmt.Println("temp--->", temp)
	return temp
}

func (withdraw *Customer) AmountWithdraw(val float64) {
	withdraw.Acc.Balance -= val
	fmt.Println("Withdraw successful!")
	fmt.Printf("The account balance after withdrawal is %.2f\n", withdraw.Acc.Balance)
	fmt.Println("----------------------------------------------------------------------------")

}

func (deposit *Customer) AmountDeposit(val float64) {
	deposit.Acc.Balance += val
	fmt.Println("Withdraw successful!")
	fmt.Printf("The account balance after withdrawal is %.2f\n", deposit.Acc.Balance)
	fmt.Println("----------------------------------------------------------------------------")
}

func (remove *Bank) CloseAccount(accountNumber int) {
	for i, customer := range remove.Cust {
		if customer.Acc.AccountNum == accountNumber {
			//remove the customer from the slice
			remove.Cust = append(remove.Cust[:i], remove.Cust[i+1:]...)
			fmt.Println("Account closed successfully.")
			fmt.Println("----------------------------------------------------------------------------")

			return
		}
	}
	fmt.Println("Account not found.")
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
		var accno int

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
			fmt.Scanln(&accno)
			selectedCustomer := FindCustomer(bank.Cust, accno)

			for {
				var amount float64

				//for to iterate to use the following options
				fmt.Println("Choose from the following options:")
				fmt.Println("1. Withdraw money")
				fmt.Println("2. Deposit money")
				fmt.Println("3. Print Balance")
				fmt.Println("4. Close the account")
				fmt.Println("5. Exit")

				var option int
				fmt.Scanln(&option)

				switch option {
				case 1:
					fmt.Println("Enter the amount you want to withdraw:")
					fmt.Scan(&amount)
					selectedCustomer.AmountWithdraw(amount)

				case 2:
					fmt.Println("Enter the amount you want to deposit:")
					fmt.Scan(&amount)
					selectedCustomer.AmountDeposit(amount)

				case 3:
					fmt.Println("The bank balance is ", selectedCustomer.Acc.Balance)
					fmt.Println("----------------------------------------------------------------------------")

				case 4:
					bank.CloseAccount(accno)

				case 5:
					fmt.Println("----------------------------------------------------------------------------")
					break

				}
				if option == 5 {
					break
				}
			}

		case 4:
			fmt.Println("You are exisiting from Banking Application.")
			fmt.Println("----------------------------------------------------------------------------")
			return

		}
	}

}
