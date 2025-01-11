package main

import "fmt"

func main() {
	// int --> depends on platform, 32/64
	var i int = 20

	// autoselect int
	var autoInt = -10

	// GO have int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// uint --> depends on platform, 32/64
	var unsignedInt uint = 100500

	// GO have uint8, uint16, uint32, uint64
	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	//float32, float64
	var pi float32 = 3.141
	var e = 2.71828
	goldenRatio := 1.618
	fmt.Println(pi, e, goldenRatio)

	// bool
	var b bool // default value is --> false
	var isOk bool = true
	var success = true
	cond := true

	fmt.Println(b, isOk, success, cond)

	// comlpex64, complex128
	var c1 complex128 = -1.1 + 7.12i
	c2 := -1.1 + 7.12i
	fmt.Println(c1, c2)

}
