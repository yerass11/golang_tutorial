package main

import "fmt"

func main() {
	var buff0 []int             // len=0, cap=0
	buff1 := []int{}            // len=0, cap=0
	buff2 := []int{44}          // len=1, cap=1
	buff3 := make([]int, 0)     // len=0, cap=0
	buff4 := make([]int, 5)     // len=5, cap=5
	buff5 := make([]int, 5, 10) // len=5, cap=10

	println(buff0, buff1, buff2, buff3, buff4, buff5)

	someInt := buff2[0]
	//someInt2 := buff2[1] -> error, because index out of range
	println(someInt)

	var buf []int
	buf = append(buf, 9, 10, 11) // len=3, cap=3
	buf = append(buf, 12)        // len=4, cap=6 -> cap x2

	otherBuf := make([]int, 3)     // [0,0,0]
	buf = append(buf, otherBuf...) // otherBuf->we gonna add other slice, otherBuf...=>we gonna add int
	fmt.Println(buf)

	var bufLen, bufCap int = len(buf), cap(buf)
	fmt.Println(bufLen, bufCap)

}
