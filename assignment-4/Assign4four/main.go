// //write a go code for, given a string and a character. Find the middle index of the given character from the string
// // Ex sample string - “this is a sample string”
// // Given character - “S”
// // Since s occurs 4 times in the above string return the 2nd index of S in the sample string which is 6
// // If the string contains the character odd number of times ex 3 times then return the floor value of the number you get after dividing it by 2 ex 3/2 = 1.5 floor of which would be 1

package main

import "fmt"

func repeat(stri string, char string) {

	var count int

	for i, _ := range stri {
		if string(stri[i]) == char {
			count++
		}
	}
	fmt.Println("first count-->", count)
	count /= 2
	fmt.Println("count-->", count)
	var tc int

	for i, _ := range stri {

		if string(stri[i]) == char {

			tc++
			if tc == count {
				fmt.Println(i)
				break
			}

		}

	}

}

func main() {
	text := "Thi is sa sample tring"
	char := "s"
	repeat(text, char)

}
