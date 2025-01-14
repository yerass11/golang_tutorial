package main

import "fmt"

func main() {
	defer fmt.Println("After work")
	fmt.Println("Some useful work")
	fmt.Println("----------------")

	fmt.Println("________________")
}
