package calculator

import "fmt"

func GetTwoInput() (int, int) {
	var int1 int
	var int2 int

	fmt.Println("Enter the first number:")
	fmt.Scanln(&int1)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&int2)

	return int1, int2
}
