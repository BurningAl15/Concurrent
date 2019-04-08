package main

import "fmt"

func main() {
	//%d for decimal, %b for binary
	fmt.Printf("%d - %b \n", 42, 42)
	//%#X for hexadecimal
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	//Loop
	for i := 100000; i < 100100; i++ {
		fmt.Printf("%d \t %b \t %#X \n", i, i, i)
	}
}
