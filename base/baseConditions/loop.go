package main

import (
	"fmt"
)

func main() {
	// Smth like while trye OR for(;;;)
	for {
		fmt.Println("loop iteration!")
		break
	}

	isRun := true
	for isRun {
		fmt.Println("loop iteration with condition!")
		isRun = false
	}

	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration with condition and index! Index: ", i)
		if i == 1 {
			continue
		}
	}

	s1 := []int{1, 2, 3}
	idx := 0
	for idx < len(s1) {
		fmt.Println("while style loop! Index: ", idx, " value: ", s1[idx])
		idx++
	}

	for idx := 0; idx < len(s1); idx++ {
		fmt.Println("c style loop! Index: ", idx, " value: ", s1[idx])
	}

	for idx := range s1 {
		fmt.Println("range slice by index:", idx, s1[idx])
	}

	for idx, val := range s1 {
		fmt.Println("range slice by idx-value", idx, val)
	}

	profile := map[int]string{1: "a", 2: "b", 3: "c"}

	for key := range profile {
		fmt.Println("range map by key:", key, profile[key])
	}

	for key, val := range profile {
		fmt.Println("range map by key-value", key, val)
	}

	for _, val := range profile {
		fmt.Println("range map by value", val)
	}

	str := "hello world"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
