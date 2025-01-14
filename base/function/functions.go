package main

import "fmt"

func singleIn(in int) int {
	return in
}

func multIn(a, b, c int) int {
	return a + b + c
}

func namedReturn() (out int) {
	out = 2
	return // by default `return out`
}

func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("%d is greater than %d", in, 2)
	}
	return in, nil
}

func multipleNamedReturn(ok bool) (rez int, err error) {
	if ok {
		err = fmt.Errorf("something went wrong")
		return
	}
	rez = 2
	return
}

func sum(in ...int) (sum int) {
	fmt.Printf("in := %#v \n", in)
	for _, v := range in {
		sum += v
	}
	return
}

func main() {
	fmt.Println(multipleNamedReturn(true))
	fmt.Println(multipleNamedReturn(false))

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(slice...))
}
