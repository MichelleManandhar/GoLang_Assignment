package main

import "fmt"

func calculation(str []string) map[string]int {
	temp := make(map[string]int)
	count := make(map[string]int)

	for _, val := range str {
		count[val]++
	}
	for key, val := range count {
		if val > 1 {
			temp[key] = val
		}
	}
	return temp

}

func main() {
	data := []string{"a", "b", "a", "b", "c", "b"}
	pass := calculation(data)
	for key, val := range pass {
		fmt.Printf("key %s is repeated %d \n", key, val)
	}

}
