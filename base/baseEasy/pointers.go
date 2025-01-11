package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 3 // a = 3
	c := &a
	fmt.Println(a, *b, *c)

	d := new(int)
	*d = 12
	*c = *d
	*d = 13 // c and a doesn't change, cause c points to a, not d
	fmt.Println(a, *b, *c, *d)

	c = d // c have pointer d
	*c = 15
	fmt.Println(a, *b, *c, *d)
}
