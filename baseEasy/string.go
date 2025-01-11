package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string

	var hello string = "Hello, world!\n\t"

	var world = `Привет, мир!\n\t` // also go have UTF-8

	fmt.Println(str, hello, world)

	//uint8
	var rawBinary byte = '\x26'

	fmt.Println(rawBinary)

	helloWorld := "hello world"
	andGoodMorning := helloWorld + " and good morning"

	// ERROR:::Cannot assign to helloWorld[0]
	//helloWorld[0] = 72

	fmt.Println(andGoodMorning)

	byteLen := len("Привет, мир!")
	symbols := utf8.RuneCountInString("Привет, мир!") // in English doesn't matter, in Russian is
	fmt.Println(byteLen, symbols)

	someString := andGoodMorning[:12] // from 0 to 11
	H := andGoodMorning[0]            // byte, 104, not "h"
	fmt.Println(someString, H)

	byteString := []byte(andGoodMorning)
	fmt.Println(byteString, andGoodMorning)
	andGoodMorning = string(byteString)

}
