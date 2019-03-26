package main

import (
	"fmt"
	"math"
	"strconv"
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

}
