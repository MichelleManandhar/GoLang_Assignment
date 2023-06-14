// 3. Write a go code to find if a given string is a palindrome or not.
// Your function should return a Boolean value telling if a given string is palindrome or not
package main

import "fmt"

func palindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	val := "abcba"
	result := palindrome(val)
	fmt.Println(result)
}
