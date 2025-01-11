package main

import "fmt"

func main() {
	var user map[string]string = map[string]string{
		"name":     "Tony",
		"lastName": "Stark",
		"midName":  "Govard",
		"age":      "42",
	}
	// map[string]string{"age":"42", "lastName":"Stark", "middleName":"Govard", "name":"Tony"}

	profile := make(map[string]string, 10)
	fmt.Println(profile)

	mapLength := len(user)
	fmt.Println(mapLength)

	mName := user["middleName"]
	fmt.Println("mName: ", mName)

	//  - map[key] в Go возвращает два значения:
	// 		* Значение, связанное с ключом (если оно существует).
	// 		* Булевое значение true или false, указывающее, существует ли ключ.
	mName, mNameExists := user["middleName"]
	fmt.Println("mName: ", mName, ", mNameExists: ", mNameExists)

	_, mNameExists2 := user["middleName"]
	fmt.Println("mNameExists2: ", mNameExists2)

	delete(user, "msda")
	fmt.Printf("%#v\n", user)
	// in alphabet order, map[string]string{"age":"42", "lastName":"Stark", "middleName":"Govard", "name":"Tony"}
}
