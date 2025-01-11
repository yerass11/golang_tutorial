package main

import "fmt"

const pi = 3.14
const (
	hello = "Heey"
	e     = 2.71828
)
const (
	zero = iota
	_
	three
)
const (
	_         = iota
	KB uint64 = 1 << (10 * iota)
	MB
)
const (
	year          = 2017
	yearTyped int = 2018
)

func main() {
	var month int32 = 13
	fmt.Println(month + year)

	// month + yearTyped -> mismatched typed int32 and int
}
