// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// // Item represents a product in the inventory
// type Item struct {
// 	ID    int
// 	Name  string
// 	Price float64
// }

// // Inventory represents the collection of available items
// type Inventory struct {
// 	Items []Item
// }

// // Cart represents a user's shopping cart
// type Cart struct {
// 	Items []Item
// }

// // ECommerceApp represents the e-commerce application
// type ECommerceApp struct {
// 	Inventory Inventory
// 	Cart      Cart
// }

// // Initialize the inventory by manually entering items
// func (app *ECommerceApp) InitializeInventory() {
// 	fmt.Println("Enter inventory items (ID, Name, Price):")
// 	for {
// 		var id int
// 		var name string
// 		var price float64

// 		fmt.Print("ID (0 to exit): ")
// 		fmt.Scanln(&id)

// 		if id == 0 {
// 			break
// 		}

// 		fmt.Print("Name: ")
// 		fmt.Scanln(&name)

// 		fmt.Print("Price: ")
// 		fmt.Scanln(&price)

// 		item := Item{
// 			ID:    id,
// 			Name:  name,
// 			Price: price,
// 		}

// 		app.Inventory.Items = append(app.Inventory.Items, item)
// 	}
// }

// // List all items in the inventory
// func (app *ECommerceApp) ListItems() {
// 	fmt.Println("Available Items:")
// 	for _, item := range app.Inventory.Items {
// 		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", item.ID, item.Name, item.Price)
// 	}
// }

// // Search for items matching the given pattern using regular expressions
// func (app *ECommerceApp) SearchItems(pattern string) {
// 	fmt.Println("Search Results:")
// 	for _, item := range app.Inventory.Items {
// 		matched, _ := regexp.MatchString(pattern, item.Name)
// 		if matched {
// 			fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", item.ID, item.Name, item.Price)
// 		}
// 	}
// }

// // Add an item to the cart
// func (app *ECommerceApp) AddToCart(itemID int) {
// 	for _, item := range app.Inventory.Items {
// 		if item.ID == itemID {
// 			app.Cart.Items = append(app.Cart.Items, item)
// 			fmt.Printf("Added '%s' to the cart.\n", item.Name)
// 			return
// 		}
// 	}
// 	fmt.Println("Item not found.")
// }

// // Calculate the total bill for the cart purchases
// func (app *ECommerceApp) CalculateTotalBill() float64 {
// 	total := 0.0
// 	for _, item := range app.Cart.Items {
// 		total += item.Price
// 	}
// 	return total
// }

// func main() {
// 	app := ECommerceApp{}
// 	app.InitializeInventory()

// 	app.ListItems()

// 	var pattern string
// 	fmt.Print("Enter search pattern: ")
// 	fmt.Scanln(&pattern)
// 	app.SearchItems(pattern)

// 	for {
// 		var itemID int
// 		fmt.Print("Enter item ID to add to cart (0 to exit): ")
// 		fmt.Scanln(&itemID)

// 		if itemID == 0 {
// 			break
// 		}

// 		app.AddToCart(itemID)
// 	}

// 	totalBill := app.CalculateTotalBill()
// 	fmt.Printf("Total Bill: %.2f\n", totalBill)
// }
