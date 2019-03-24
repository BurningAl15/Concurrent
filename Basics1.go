package main

import (
	"fmt"
	"math"

	//Tutorial_01
	"time"
	//Tutorial_02
	"math/cmplx"
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

//Tutorial 8
func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x
	z = sum - x - (y / 2)
	return
}

//Tutorial_12
var (
	ToBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var cSharp, python, java bool

var i, j int = 1, 2

const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

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
	fmt.Println(i, j, cSharp, python, java)

	//Tutorial_11
	//Short variable declarations
	TutorialText("11", "Short Var declarations")
	k := 3
	fmt.Println(i, j, k, cSharp, python, java)

	//Tutorial_12
	//Basic types
	TutorialText("12", "Basic Types")
	/*
		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte //Alias for uint8

		rune //alias for int32
			 //represents a Unicode code point

		float32 float64

		complex64 complex128
	*/
	fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T value: %v\n", z, z)

	//Tutorial_13
	//Zero Values
	TutorialText("13", "Zero Values")
	/*
		0 for numeric types,
		false for boolean types
		"" (empty string) for strings
	*/

	var a int
	var b float64
	var d bool
	var e string
	fmt.Printf("%v %v %v %q\n", a, b, d, e)

	//Tutorial_14
	//Type Conversion
	TutorialText("14", "Type Casting")
	i = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)

	//Tutorial_15
	//Type inference
	TutorialText("15", "Type Inference")
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	//Tutorial_16
	//Constants
	TutorialText("16", "Constants")

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	//Tutorial_17
	//Numeric Constants
	TutorialText("17", "Constants")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
