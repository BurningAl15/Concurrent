package main

import "fmt"

type Loud interface {
	beLoud(str string) string
}

type Dog struct {
	name string
}

type Singer struct {
	bandName string
}

func (fido Dog) beLoud(str string) string {
	return fmt.Sprintf("%s: GUAU GUAU GUAU GUAU (%s)", fido.name, str)
}

func (fido Singer) beLoud(str string) string {
	return fmt.Sprintf("%s: Hey Yo we are the Jons (%s)", fido.bandName, str)
}

func doStuff(a Loud, msg string) {
	fmt.Println(a.beLoud(msg))
}

func main() {
	fido := Dog{"Fido"}
	jon := Singer{"The Jons"}

	doStuff(fido, "Hola")
	doStuff(jon, "Hey Yo")
}
