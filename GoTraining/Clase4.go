package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func Case1InitFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func Case2NotInitFor() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Case3ForGoWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Case4InfiniteLoop() {
	i := 0
	for {
		fmt.Println("A" + string(i))
		i++
	}
}

func Case5Sqrt(x float64) string {
	if x < 0 {
		return Case5Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func Case6Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Case7Switch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		//freebsd,openbsd
		//plan9,windows
		fmt.Printf("%s. \n", os)
	}
}

func Case8() {
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
}

func Case9KnowTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func Case10Defer() {
	defer fmt.Println("!")
	defer fmt.Println("world")
	fmt.Println("Hello")
}

func Case11StackDefer() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func Case12Pointers() {
	i, j := 42, 2701
	//& -> point to
	//* -> read the pointer

	p := &i         //point to i
	fmt.Println(*p) //read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)  //see the new value of i

	p = &j         //point to j
	*p = *p / 37   //divide j through the pointer
	fmt.Println(j) //see the new value of j
}

type Vertex struct {
	X int
	Y int
}

func Case1_b_Structs() {
	fmt.Println(Vertex{1, 2})
}

func Case2_b_Structs(_a Vertex, index int) {
	switch index {
	case 0:
		fmt.Println(_a.X)
	case 1:
		fmt.Println(_a.Y)
	default:
		fmt.Println(_a)
	}
}

func Case3_b_StructsChangeVal(_a Vertex, index int, val int) {
	switch index {
	case 0:
		_a.X = val
		fmt.Println(_a.X)
	case 1:
		_a.Y = val
		fmt.Println(_a.Y)
	default:
		fmt.Println(_a)
	}
}

var (
	v1 = Vertex{1, 2}  //has type Vertex
	v2 = Vertex{X: 1}  //Y:0 is implicit
	v3 = Vertex{}      //X:0 and Y:0
	p  = &Vertex{1, 2} //has type *Vertex
)

func Case1_Array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func Case2_Array() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func Case1_Slices(from int, to int) {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[from:to]
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	//Case1InitFor()
	//Case2NotInitFor()
	//Case3ForGoWhile()
	//Case4InfiniteLoop()

	//fmt.Println(Case5Sqrt(2), Case5Sqrt(-4))

	// fmt.Println(
	// 	Case6Pow(3, 2, 10),
	// 	Case6Pow(3, 3, 20),
	// )

	//Case7Switch()
	//Case8()
	//Case9KnowTime()
	//Case10Defer()

	//Case11StackDefer()
	//Case12Pointers()

	// a := Vertex{1, 2}

	// Case2_b_Structs(a, 0)
	// Case2_b_Structs(a, 1)
	// Case2_b_Structs(a, 2)

	// Case3_b_StructsChangeVal(a, 0, 2)
	// Case3_b_StructsChangeVal(a, 1, 5)
	// Case3_b_StructsChangeVal(a, 2, 5)

	// p := &a
	// p.X = 1e9
	// fmt.Println(a)
	// fmt.Println(v1, p, v2, v3)
	// Case2_Array()

	// Case1_Slices(0, 6)
	// Case1_Slices(1, 4)

	// var s []int
	// printSlice(s)

	// // append works on nil slices.
	// s = append(s, 0)
	// printSlice(s)

	// // The slice grows as needed.
	// s = append(s, 1)
	// printSlice(s)

	// // We can add more than one element at a time.
	// s = append(s, 2, 3, 4)
	// printSlice(s)

	// go say("world")
	// say("hello")

	// c := SafeCounter{v: make(map[string]int)}
	// for i := 0; i < 1000; i++ {
	// 	go c.Inc("somekey")
	// }
	// time.Sleep(time.Second)
	// fmt.Println(c.Value("somekey"))
}
