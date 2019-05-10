package main

import (
	"fmt"
	"math/rand"
	"time"
)

var wantP = false
var wantQ = false
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

func p() {
	// for int i=0;i<10;i++ {
	for {
		Pause("P Msg 1:No critical section")
		wantP = true
		for !wantQ {
			Pause("P Msg 1:Critical section X")
		}
		wantQ = false
		end <- true
	}
}

func q() {
	// for int i=0;i<10;i++ {
	for {
		Pause("P Msg 1:No critical section")
		wantQ = true
		for !wantP {
			Pause("P Msg 1:Critical section X")
		}
		wantP = false
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
