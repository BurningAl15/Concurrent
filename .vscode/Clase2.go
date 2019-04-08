package main

import (
	"fmt"
	"time"
)

var n int = 0

func P() {
	var temp int
	for i := 0; i < 10; i++ {
		Pause()
		temp = n
		Pause()
		n = temp + 1
	}
	// fmt.Print(n)
}

func Q() {
	var temp int
	for i := 0; i < 10; i++ {
		Pause()
		temp = n
		Pause()
		n = temp + 1
	}
	// fmt.Print(n)
}

func Pause() {
	time.Sleep(100 * time.Millisecond)
}

type Animal struct {
	name string
}

func main() {
	// rand.
	// go P()
	// go Q()
	// time.Sleep(30 * time.Millisecond)
	// fmt.Println(n)
	a := 10
	ptr := &a

	fmt.Println(a, ptr, *ptr)
	var ptr2 *int
	ptr2 = new(int)
	*ptr2 = 20
	fmt.Println(ptr2, *ptr2)

	var a1 Animal
	a1.name = "Fido"
	a2 := Animal{"Spike"}
	a3 := Animal{name: "Spike"}
	ptr3 := &a3
	ptr4 := &Animal{"Michifuss"}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(ptr3, *ptr3)
	fmt.Println(ptr4, *ptr4)
}
