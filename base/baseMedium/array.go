package main

import "fmt"

func main() {
	var a1 [3]int // [0, 0, 0]
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	const size = 2        // we can use ```const``` for arr size, but can't ```int```
	var a2 [2 * size]bool // [false false false false]
	fmt.Println("a2", a2)

	a3 := [...]int{1, 2, 3}
	fmt.Println("a3", a3)

	a3[2] = 11
	//a3[3] = 12 // : invalid argument: index 3 out of bounds [0:3]
	fmt.Println("a3", a3)
}
