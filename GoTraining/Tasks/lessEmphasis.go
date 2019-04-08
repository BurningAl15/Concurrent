package main

import "fmt"

//All together
var a = "this is stored in the variable a"
var b, c string = "stored in b", "stored in c"
var d string

func main() {
	//Declaration way 2
	// var message string = "Hello world."
	// var a, b, c int = 1, 2, 3

	//Declaration way 1
	// a = 1
	// message = "Hello world."

	//Infer type
	// var message = "Hello World!"
	// var a, b, c = 1, 2, 3

	//Infer mixed types
	// var message = "Hello World!"
	// var a, b, c = 1, false, 3

	//Init shorthand
	// message := "Hello world"
	// a, b, c := 1, false, 3
	// d := 4
	// e := true

	// fmt.Println(message, a, b, c, d, e)

	d = "stored in d"
	var e = 42
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm'
	n := "n"
	o := `o`
	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
}
