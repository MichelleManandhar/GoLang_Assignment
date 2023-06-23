package main

import (
	"fmt"
	"time"
)

func main() {

	time.AfterFunc(time.Second*6, func() {
		fmt.Println("first")
	})

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("second")

	}()

	time.Sleep(time.Second * 9)
	fmt.Println("third")

}
