package main

import (
	"fmt"
	"math"

	//Tutorial_01
	"time"
	//Tutorial_02
	"math/rand"
)

func TutorialText(tutorialName string, theme string) {
	fmt.Println(" ")
	fmt.Println("/////////////////////////////////////")
	fmt.Printf("///////       Tutorial %s     /////// \n", tutorialName)
	fmt.Printf("///////    Theme:  %s  /////// \n", theme)
	fmt.Println("/////////////////////////////////////")
	fmt.Println(" ")
}

//Tutorial 5
// func add(x int, y int) int {
// 	return x + y
// }

//Tutorial 6
func add(x, y int) int {
	return x + y
}

//Tutorial 7
func swap(x, y string) (string, string) {
	return y, x
}

//Tutorial 7
func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x
	z = sum - x - (y / 2)
	return
}

var cSharp, python, java bool

func main() {
	//fmt.Println("Hello, 世界")

	//Tutorial_01
	//How to print stuff and use time import module
	TutorialText("01", "Println")
	fmt.Println("The time is", time.Now())

	//Tutorial_02
	//How to Print random numbers using time as seed
	TutorialText("02", "Random")
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is: ", rand.Intn(10))

	//Tutorial_03
	//How to use format for float values
	TutorialText("03", "Formats")
	//In case u use fmt.Println, the value will appear in the next line
	//Mean %g is evaded
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))

	//Tutorial_04
	//Exported Name
	//math.Pi, Pi is an exported name which is exported
	//from the math package
	TutorialText("04", "Exported Name")
	fmt.Println(math.Pi)

	//Tutorial_05
	//Functions
	TutorialText("05", "Functions")
	fmt.Println(add(101, 99))

	//Tutorial_06
	//Functions
	TutorialText("06", "Functions continued")
	//x int, y int are the same as x,y int as parameters
	//From a function
	fmt.Println(add(101, 99))

	//Tutorial_07
	//Multiple Results
	TutorialText("07", "Multiple results")
	//Assign values
	c, x := swap("Hello", "World")
	fmt.Println(x, c)
	fmt.Println(c, x)

	//Tutorial_08
	//Named return values
	TutorialText("08", "Named Return values")
	fmt.Println(split(20))

	//Tutorial_09
	//Variables
	TutorialText("09", "Variables")
	var i int
	fmt.Println(i, cSharp, python, java)

	//Tutorial_10
	//Variables with initializers
	TutorialText("10", "Var with inits")
	TutorialText("10", "Var with inits")

}
