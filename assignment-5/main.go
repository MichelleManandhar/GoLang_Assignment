// Banking Application
package main

import "fmt"

type Bank struct {
	//field called Customers which has slie of ustomer object
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

func (numcheck *Bank) FindCustomer(accno int) Customer {

	var val Customer
	for _, value := range numcheck.Cust {
		//Output 1
		//fmt.Println("numcheck->", numcheck, "  value->", value)
		//numcheck-> &{[{miche {123 3434}} {tha {456 454545}}]}   value-> {miche {123 3434}}
		//numcheck-> &{[{miche {123 3434}} {tha {456 454545}}]}   value-> {tha {456 454545}}
		//Output 2
		// fmt.Println("index->", index, "  value->", value)
		// index-> 0   value-> {mich {123 343434}}
		// index-> 1   value-> {tha {456 454545}}
		if value.Acc.AccountNum == accno {
			// fmt.Println("####---->", value.Acc.AccountNum)  output : 123
			// fmt.Println("index->", index, " value->", value)
			val = value
			fmt.Println("You have accessed the details of customer = ", value.Name)
			fmt.Println("----------------------------------------------------------------------------")

		}
		// else {
		// 	fmt.Println("Account Number you entered is invalid.")
		// 	fmt.Println("----------------------------------------------------------------------------")
	}
	fmt.Println("val---->", val)
	//output : val----> {tha {2 3.2423432e+07}}
	return val
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
			bank.FindCustomer(accno)

		case 4:
			fmt.Println("You are exisiting from Banking Application.")
			fmt.Println("----------------------------------------------------------------------------")
			return

		}
	}

}
