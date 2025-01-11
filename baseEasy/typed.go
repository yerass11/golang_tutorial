package main

import "fmt"

type UserID int

func main() {
	idx := 1
	var uid UserID = 42
	myID := idx
	fmt.Println(idx, uid, myID)

	myID2 := UserID(idx)
	println(uid, myID, myID2)
}
