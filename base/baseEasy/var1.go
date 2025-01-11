package main

import "fmt"

func main() {
	var num0 int // var varName varType --> default value is 0

	var num1 int = 1 // var varName varType = varValue

	var num3 = 2 // var varName = varValue --> type inference

	fmt.Println(num0, num1, num3)

	num := 30 // new var, type inference
	//num := 31  -->  No new variables on the left side of ':=', it means if var is exists, we can't use it like that
	num += 1
	fmt.Println("+=", num)

	num++ // u can use num++, --> but ++num ERROR
	fmt.Println("++", num)

	userIndex := 10  // good
	user_index := 20 // --> not ERROR, but not good
	fmt.Println(userIndex, user_index)

	var weight, height = 90, 180 // var varName1, varName2 = varValue1, varValue2
	fmt.Println(weight, height)

	weight, height = 83, 180 // varName1, varName2 = varValue1, varValue2

	//weight, height := 90, 180 -> You need at least 1 variable left side of :=

	weight, age := 86, 20 // -> age is new variable, so we can use weight
	fmt.Println("weight:", weight, "\nheight:", height, "\nage:", age)

}
