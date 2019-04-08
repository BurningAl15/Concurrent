package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"
)

func TutorialText(tutorialName string, theme string) {
	fmt.Println(" ")
	fmt.Println("/////////////////////////////////////")
	fmt.Printf("///////       Tutorial %s     /////// \n", tutorialName)
	fmt.Printf("///////    Theme:  %s  /////// \n", theme)
	fmt.Println("/////////////////////////////////////")
	fmt.Println(" ")
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	//For loop
	sum := 0
	for i := 0; i < 25; i++ {
		toString := strconv.Itoa(i)
		if i < 10 {
			TutorialText("0"+toString, "Theme "+toString)
		} else {
			TutorialText(toString, "Theme "+toString)
		}

		sum += i
	}
	fmt.Println(sum)
	fmt.Println(" ")

	//Continued For loop
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	fmt.Println(" ")

	//Infinite loop
	i := 0
	for {
		i++
		fmt.Println(i)
		if i > 1000 {
			break
		}
	}

	fmt.Println(" ")

	//Short if
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(" ")

	//Switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//freebsd, openbsd
		//plan9,windows ...
		fmt.Printf("%s. \n", os)
	}
	fmt.Println(" ")

	//
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Println(" ")

	//
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println(" ")

	//Defer
	defer fmt.Println("world!")
	fmt.Println("hello")
	fmt.Println(" ")

	//Stacking Defer
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	fmt.Println(" ")
}
