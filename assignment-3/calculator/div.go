package calculator

import "fmt"

func Div(int1, int2 int) {
	quo := (int1 / int2)
	rem := (int1 % int2)
	fmt.Println("The quotient of the two numbers is:", quo)
	fmt.Println("The remainder of the two numbers is:", rem)
}
