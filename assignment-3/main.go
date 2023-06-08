package main

import (
	"GoLang_Assignment/assignment-3/calculator"
)

func main() {
	first, second := calculator.GetTwoInput()
	//fmt.Println("first , sec is", first, second)
	calculator.Add(first, second)
	calculator.Sub(first, second)
	calculator.Mul(first, second)
	calculator.Div(first, second)
}
