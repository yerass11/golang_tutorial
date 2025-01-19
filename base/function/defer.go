package main

import "fmt"

func getSomeVars() string {
	fmt.Println("getSomeVars println")
	return "getSomeVars return"
}

func main() {
	defer fmt.Println("After work")
	// defer fmt.Println(getSomeVars())
	defer func() {
		fmt.Println(getSomeVars())
	}()
	fmt.Println("Some useful work")
	fmt.Println("________________")

}
