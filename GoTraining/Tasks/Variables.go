package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat`
	g := 'M'
	//%v looks like gives the value
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Println(" ")
	//%T gives the type
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
	fmt.Println()

	//////////////
	// Var zero //
	//////////////
	var h int
	var i string
	var j float64
	var k bool
	fmt.Printf("int: %v \n", h)
	fmt.Printf("string: %v \n", i)
	fmt.Printf("float64: %v \n", j)
	fmt.Printf("bool: %v \n", k)
	fmt.Println()
}
