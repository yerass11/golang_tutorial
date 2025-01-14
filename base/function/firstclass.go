package main

import "fmt"

func doNothing() {
	fmt.Println("i am regular function")
}

func main() {
	func(in string) {
		fmt.Println("anon func out :", in)
	}("nobody")

	printer := func(in string) {
		fmt.Println("printer out :", in)
	}
	printer("as variable")

	type strFuncType func(in string)

	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(printer)

	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}

	succesLogger := prefixer("SUCCESS")
	succesLogger("expected behaviour")
	return

	fmt.Println("i am regular function")
}
