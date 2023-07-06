package main

import (
	"fmt"
	"regexp"
)

type Ecommerce struct {
	Inv Inventory
}

type Items struct {
	ID      int
	Product string
	Price   float64
}

type Inventory struct {
	Item []Items
	Cart []Carts
}

type Carts struct {
	ID      int
	Product string
	Price   float64
}

func (insert *Inventory) InsertItems() {
	fmt.Println("Add items in the inventory (ID, Name, and Price) or press 0 to exit")
	for {
		var id int
		var product string
		var price float64

		fmt.Println("Insert the ID or press 0 to exit")
		fmt.Scanln(&id)
		if id == 0 {
			fmt.Println("******Existing******")
			break
		}

		fmt.Println("Insert the product name:")
		fmt.Scanln(&product)

		fmt.Println("Insert the price of the product:")
		fmt.Scanln(&price)

		list := Items{
			ID:      id,
			Product: product,
			Price:   price,
		}
		insert.Item = append(insert.Item, list)
		fmt.Println("----------------------------------------------------------------------------")
	}
	fmt.Println("----------------------------------------------------------------------------")
}

func (view *Inventory) DisplayItems() {
	fmt.Println("The available items in the list are:")
	for _, value := range view.Item {
		fmt.Printf("ID: %d, Product: %s, Price: %.2f\n", value.ID, value.Product, value.Price)
	}
	fmt.Println("----------------------------------------------------------------------------")
}

func (search *Inventory) SearchItems(pattern string) {
	//search based on regexp example : ^A.*, .*shirt.*, .*s$
	fmt.Println("The search result:")
	for _, value := range search.Item {
		matched, _ := regexp.MatchString(pattern, value.Product)
		// search inventory item--> 0   value {1 tamato 2}
		// search inventory item--> 1   value {2 potato 3}
		// fmt.Println("search inventory item-->", item, "  value", value)
		if matched == true {
			fmt.Printf("ID: %d, Product: %s, Price: %.2f\n", value.ID, value.Product, value.Price)
		}
	}
}

func (val *Inventory) AddToCart(itemId int) {
	//adding selected item to cart where the user inputs itemId
	for _, item := range val.Item {
		// fmt.Println("cart item-->", a, "  value", item)
		if item.ID == itemId {
			val.Cart = append(val.Cart, Carts(item))
			fmt.Printf("Added %s to the cart.\n", item.Product)
			return
		}
	}
	fmt.Println("Item not found.")
}

func (bill *Inventory) TotalBill() float64 {
	//adding up the total that was added to the cart
	var amount float64
	for _, item := range bill.Cart {
		amount += item.Price
	}
	return amount
}

func main() {
	inv := Inventory{}
	fmt.Println("**********Welcome to the E-commerce Application**********")
	//Function to insert the items
	inv.InsertItems()
	// fmt.Println(inv.Item)
	//Functin to print the items in the slice
	inv.DisplayItems()

	var pattern string
	fmt.Print("Enter the search pattern:")
	fmt.Scanln(&pattern)
	inv.SearchItems(pattern)

	for {
		var itemId int
		fmt.Print("Enter the item ID to add to cart or press 0 to exit:")
		fmt.Scanln(&itemId)
		//id selected by the customer to add in the cart
		if itemId == 0 {
			break
		}
		inv.AddToCart(itemId)
	}
	//printing the total which is in the cart
	totalAmt := inv.TotalBill()
	fmt.Println("The total bill is $", totalAmt)
}
