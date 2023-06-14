// 1. write a go code to find even and odd elements of a slice.
// Your answer should return 2 slices containing even and odd elements correspondingly
package main

import "fmt"

func calculation(a []int) ([]int, []int) {
	even := []int{}
	odd := []int{}
	for i := range a {
		if a[i]%2 == 0 {
			even = append(even, a[i])
		} else {
			odd = append(odd, a[i])
		}
	}
	return even, odd
}

func main() {
	number := []int{11, 12, 3, 4, 5, 16, 17, 8, 9, 10}
	even1, odd1 := calculation(number)
	//even1=checkiflessthan10(even1)
	//odd1=checkiflessthan10(odd1)

	fmt.Println("The even slices is=", even1)
	fmt.Println("The odd slices is=", odd1)
}

// func checkiflessthan10(a []int) []int {

// 	less10 := []int{}

// 	for i := range a {
// 		if a[i] < 10 {
// 			less10 = append(less10, a[i])

// 		}

// 	}
// 	return less10
// }
