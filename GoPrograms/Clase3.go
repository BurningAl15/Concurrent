package main

import (
	"fmt"
	"math/rand"
	"time"
)

var turn int = 1

var end chan bool

func Pause(message string) {
	dur := time.Duration(rand.Intn(100) + 50)
	time.Sleep(dur * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("/////////////////////////////////////")
	fmt.Printf("       %s      \n", message)
	fmt.Println("/////////////////////////////////////")
	fmt.Println(" ")
}

func AssignTurn(_turn int) {
	turn = _turn
	fmt.Println(" ")
	fmt.Printf("///////       %d     /////// \n", turn)
	fmt.Println(" ")
}

func p() {
	// for int i=0;i<10;i++ {
	for {
		Pause("P Msg 1:No critical section")
		Pause("P Msg 2:No critical section")
		Pause("P Msg 3:No critical section")
		for turn != 2 {

		}
		Pause("P Msg 1:Critical section X")
		Pause("P Msg 2:Critical section X")
		Pause("P Msg 3:Critical section X")
		turn = 2
		end <- true
	}
}

func q() {
	// for int i=0;i<10;i++ {
	for {
		Pause("Q Msg 1:No critical section")
		Pause("Q Msg 2:No critical section")
		Pause("Q Msg 3:No critical section")
		for turn != 1 {

		}
		Pause("Q Msg 1:Critical section X")
		Pause("Q Msg 2:Critical section X")
		Pause("Q Msg 3:Critical section X")
		turn = 1
		end <- true
	}
}

func main() {
	end = make(chan bool)
	go p()
	go q()
	<-end
	<-end
	// time.Sleep(10 * time.Second)
}
