package main

import (
	"fmt"
	"regexp"
)

type Ecommerce struct {
	Inv Inventory
	Car Cart
}

type Items struct {
	ID      int
	Product string
	Price   float64
}

type Inventory struct {
	Item []Items
}

type Cart struct {
	Item []Items
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

func (cart *Ecommerce) AddToCart(id int) {
	fmt.Println("The value of the id ---->", id)
	for item, value := range cart.Inv.Item {
		if value.ID == id {
			cart.Car.Item = append(cart.Car.Item, value)
			fmt.Printf("Added %s to the cart.\n", value.Product)
		}
		fmt.Println("search inventory item-->", item, "  value", value)
	}
	fmt.Println("Exited the function")
}

func main() {
	inv := Inventory{}
	app := Ecommerce{}
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
	// 	To search for items that start with the letter "A": ^A.*
	// To search for items that contain the word "shirt": .*shirt.*
	// To search for items that end with the letter "s": .*s$

	for {
		var cartItem int
		fmt.Print("Enter the item ID to add to cart or press 0 to exit:")
		fmt.Scanln(&cartItem)
		if cartItem == 0 {
			break
		}
		fmt.Println("before")
		app.AddToCart(cartItem)
		fmt.Println("after")
	}
	fmt.Println("Outside the function")
}
