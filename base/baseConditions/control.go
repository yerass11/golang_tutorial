package main

import "fmt"

func main() {
	boolVal := true
	if boolVal {
		fmt.Println("Bool is true")
	}

	mapVal := map[string]int{"a": 1, "b": 2}
	if keyValue, keyExists := mapVal["a"]; keyExists {
		fmt.Println("Key a has value", keyValue)
	}

	if _, keyExists := mapVal["a"]; keyExists {
		fmt.Println("Key a has value")
	}

	cond := 1
	if cond == 1 {
		fmt.Println("cond is 1")
	} else {
		fmt.Println("cond is not 1")
	}

	strVal := "name"
	switch strVal {
	case "name":
		fallthrough // выполняет след кейс без проверки условия
	case "age", "weight":
		fmt.Println("strVal is name, weight or age")
	default:
		fmt.Println("strVal is not name, age or weight")
	}

	val1, val2 := 10, 15
	switch {
	case val1 > 5 || val2 > 10:
		fmt.Println("val1 > 5 || val2 > 10")
	case val2 < 8:
		fmt.Println("val1 <= 5 && val2 <= 10")
	}
Loop:
	for key, val := range mapVal {
		fmt.Println("switch in loop", key, val)
		switch {
		case key == "lastName":
			break
			//print("don't print this")
		case key == "firstName" && val == 1:
			break Loop
		}
	}
}
