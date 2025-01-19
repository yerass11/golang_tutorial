package main

import "fmt"

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend:", err)
		}
	}()
	fmt.Println("defer testing")
	panic("panic test")

	return
}

func main() {
	fmt.Println("main")

	deferTest()

	fmt.Println("after panic test")
}
