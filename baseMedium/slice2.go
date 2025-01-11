package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	sl1 := buf[1:4] // [2,3,4]
	sl2 := buf[:2]  // 1, 2
	sl3 := buf[2:]  // [3, 4, 5]
	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:]
	newBuf[0] = 100 // same memory
	fmt.Println("newBuf: ", newBuf)
	fmt.Println("buf: ", buf)

	newBuf = append(newBuf, 6) // we created new memory for newBuf, after that copy values

	newBuf[0] = 1
	fmt.Println("newBuf: ", newBuf)
	fmt.Println("buf: ", buf)

	var emptyBuf []int
	copied := copy(emptyBuf, buf) // copied = 0
	fmt.Printf("copied: %d, emptyBuf: %d\n", copied, emptyBuf)
	fmt.Println("buf: ", buf)

	// right method
	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)
	newBuf[0] = 99
	fmt.Println("newBuf: ", newBuf)
	fmt.Println("buf: ", buf)

	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6, 7}) // ints = [1, 5, 6, 4]
	fmt.Println("ints: ", ints)

}
