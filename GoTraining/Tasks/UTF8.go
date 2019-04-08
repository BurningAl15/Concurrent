package main

import "fmt"

func main() {
	//UTF-8

	fmt.Println(" ")
	value := 65

	fmt.Printf("%d \t %b \t %x \t %q \n", value, value, value, value)

	for i := 60; i <= 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
